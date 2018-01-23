package airbrake

import (
	stdlog "log"

	"github.com/airbrake/gobrake"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// SetLogger is an invokable function making gobrake to use the application logger.
func SetLogger(logger log.Logger) {
	gobrake.SetLogger(stdlog.New(log.NewStdlibAdapter(level.Error(logger)), "gobrake: ", 0))
}
