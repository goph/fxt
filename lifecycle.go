package fxt

import (
	"context"

	"go.uber.org/fx"
)

// Closer is the interface after io.Closer for post-shutdown cleanups.
type Closer interface {
	Close() error
}

// closerFunc makes a function matching Closer.Close signature a Closer instance.
type closerFunc func() error

func (fn closerFunc) Close() error {
	return fn()
}

// Hook is a set of callbacks extending fx.Hook.
// It adds an OnClose hook which allows post-shutdown cleanups.
// It's signature matches with the Closer interface.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
	OnClose func() error
}

// Lifecycle is the extended version of fx.Lifecycle.
type Lifecycle interface {
	Append(Hook)
}

// NewLifecycle creates a new lifecycle required by all constructors providing closers.
func NewLifecycle(l fx.Lifecycle) (Lifecycle, Closer) {
	lf := &lifecycle{
		lifecycle: l,
	}

	return lf, lf
}

type lifecycle struct {
	lifecycle fx.Lifecycle
	closers   []Closer
}

func (l *lifecycle) Append(hook Hook) {
	l.lifecycle.Append(fx.Hook{
		OnStart: hook.OnStart,
		OnStop:  hook.OnStop,
	})

	l.closers = append(l.closers, closerFunc(hook.OnClose))
}

func (l *lifecycle) Close() error {
	// TODO: handle multi errors
	for _, closer := range l.closers {
		closer.Close()
	}

	return nil
}
