package fxpromhttp

import (
	stdlog "log"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
)

// LoggerParams provides a set of dependencies for a HTTP handler logger constructor.
type LoggerParams struct {
	dig.In

	Logger log.Logger `optional:"true"`
}

// NewLogger creates a new, promhttp compliant logger instance.
func NewLogger(params LoggerParams) promhttp.Logger {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	return stdlog.New(fxlog.NewWriterAdapter(level.Error(log.With(logger, "component", "prometheus"))), "", 0)
}
