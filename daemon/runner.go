package daemon

import (
	"context"
	"errors"
	"sync"

	"github.com/goph/fxt"
)

// Runner runs a daemon.
type Runner struct {
	Daemon Daemon

	mu       sync.Mutex
	doneChan chan struct{}
	quitChan chan struct{}
}

// NewRunner creates a new daemon runner and registers it in the application lifecycle.
func NewRunner(daemon Daemon, l fxt.Lifecycle) (*Runner, Err) {
	runner := &Runner{
		Daemon: daemon,
	}

	errCh := make(chan error, 1)

	l.Append(fxt.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				errCh <- runner.Run()
			}()

			return nil
		},
		OnStop: runner.Shutdown,
	})

	return runner, errCh
}

func (s *Runner) getQuitChan() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.getQuitChanLocked()
}

func (s *Runner) getQuitChanLocked() chan struct{} {
	if s.doneChan == nil {
		s.doneChan = make(chan struct{})
	}
	return s.doneChan
}

func (s *Runner) closeQuitChanLocked() {
	ch := s.getDoneChanLocked()

	select {
	case <-ch:
		// Already closed. Don't close again.

	default:
		// Safe to close here. We're the only closer,
		// guarded by s.mu.
		close(ch)
	}
}

func (s *Runner) getDoneChanLocked() chan struct{} {
	if s.doneChan == nil {
		s.doneChan = make(chan struct{})
	}
	return s.doneChan
}

func (s *Runner) closeDoneChanLocked() {
	ch := s.getDoneChanLocked()

	select {
	case <-ch:
		// Already closed. Don't close again.

	default:
		// Safe to close here. We're the only closer,
		// guarded by s.mu.
		close(ch)
	}
}

// Run runs the daemon itself.
func (s *Runner) Run() error {
	if s.Daemon == nil {
		return errors.New("no daemon specified")
	}

	err := s.Daemon.Run(s.getQuitChan())

	s.mu.Lock()
	s.closeDoneChanLocked()
	s.mu.Unlock()

	return err
}

// Shutdown gracefully stops the daemon by sending a quit signal to it and waiting for it to be done.
func (s *Runner) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	s.closeQuitChanLocked()
	s.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-s.doneChan:
		return nil
	}
}

// Close sends a quit signal to the daemon.
func (s *Runner) Close() error {
	s.mu.Lock()
	s.closeQuitChanLocked()
	s.mu.Unlock()

	return nil
}
