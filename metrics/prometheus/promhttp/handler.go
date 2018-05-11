package fxpromhttp

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
)

// HandlerParams provides a set of dependencies for a HTTP handler constructor.
type HandlerParams struct {
	dig.In

	Gatherer prometheus.Gatherer

	Logger promhttp.Logger      `optional:"true"`
	Opts   promhttp.HandlerOpts `optional:"true"`
}

// Handler is a unique type intended to be the return type for the handler constructor.
// Consuming packages can use this interface the register the handler in a HTTP mux.
type Handler interface {
	http.Handler
}

// NewHandler returns a new HTTP handler for Prometheus metrics.
func NewHandler(params HandlerParams) Handler {
	opts := params.Opts
	if params.Opts.ErrorLog == nil {
		params.Opts.ErrorLog = params.Logger
	}

	return promhttp.HandlerFor(params.Gatherer, opts)
}
