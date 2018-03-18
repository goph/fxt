package prometheus

import (
	stdlog "log"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	fxlog "github.com/goph/fxt/log"
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

// NewLogger creates a new, promhttp compliant logger instance.
func NewLogger(logger log.Logger) (promhttp.Logger) {
	return stdlog.New(fxlog.NewWriterAdapter(level.Error(logger)), "prometheus", 0)
}
