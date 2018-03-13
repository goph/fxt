package prometheus

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterHandler registers the Prometheus metrics handler in the debug server.
func RegisterHandler(params HandlerParams) {
	opts := params.Opts
	if params.Opts.ErrorLog == nil {
		params.Opts.ErrorLog = params.Logger
	}

	params.Handler.Handle("/metrics", promhttp.HandlerFor(params.Gatherer, opts))
}

// promLogger is a simple implementation of the promhttp.Logger interface.
type promLogger struct {
	logger log.Logger
}

func (l *promLogger) Println(v ...interface{}) {
	l.logger.Log("msg", fmt.Sprintln(v...))
}

// NewLogger creates a new, promhttp compliant logger instance.
func NewLogger(logger log.Logger) (promhttp.Logger) {
	return &promLogger{logger}
}
