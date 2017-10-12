package daemon

import (
	"errors"
	"time"
)

// CronJob is ran by the CronDaemon.
type CronJob interface {
	// Run runs the job.
	Run() error
}

// CronDaemon is a daemon with an internal scheduler for a CronJob.
type CronDaemon struct {
	Job    CronJob
	Ticker *time.Ticker

	// Blocking means the scheduler does not allow jobs to overlap, there can only be one running at a time.
	Blocking bool
}

// Run implements the Daemon interface.
func (d *CronDaemon) Run(quit <-chan struct{}) error {
	if d.Job == nil {
		return errors.New("no job specified")
	}

	if d.Ticker == nil {
		return d.Job.Run()
	}

	if d.Blocking {
		return d.runBlocking(quit)
	}

	return d.runNonBlocking(quit)
}

func (d *CronDaemon) runBlocking(quit <-chan struct{}) error {
	for {
		select {
		case <-quit:
			return nil

		case <-d.Ticker.C:
			err := d.Job.Run()
			if err != nil {
				return err
			}
		}
	}
}

func (d *CronDaemon) runNonBlocking(quit <-chan struct{}) error {
	errChan := make(chan error)

	for {
		select {
		case <-quit:
			return nil

		case err := <-errChan:
			return err

		case <-d.Ticker.C:
			go func() {
				err := d.Job.Run()

				if err != nil {
					errChan <- err
				}
			}()
		}
	}
}
