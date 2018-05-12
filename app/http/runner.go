package fxhttpapp

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt/debug"
	"github.com/goph/fxt/http"
	"github.com/goph/healthz"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

// Runner executes the application and waits for it to end.
type Runner struct {
	fx.In

	Logger log.Logger
	Status *healthz.StatusChecker

	DebugErr fxdebug.Err
	HTTPErr  fxhttp.Err
}

// Run waits for the application to finish or exit because of some error.
func (r *Runner) Run(app interface {
	Done() <-chan os.Signal
}) error {
	defer func() {
		level.Debug(r.Logger).Log("msg", "setting application status to unhealthy")
		r.Status.SetStatus(healthz.Unhealthy)
	}()

	select {
	case sig := <-app.Done():
		fmt.Println() // empty line before log entry
		level.Info(r.Logger).Log("msg", fmt.Sprintf("captured %v signal", sig))

	case err := <-r.DebugErr:
		if err != nil && err != fxdebug.ErrServerClosed {
			return errors.Wrap(err, "debug server crashed")
		}

	case err := <-r.HTTPErr:
		if err != nil && err != fxhttp.ErrServerClosed {
			return errors.Wrap(err, "http server crashed")
		}
	}

	return nil
}
