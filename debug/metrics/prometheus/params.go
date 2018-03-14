package prometheus

import (
	"github.com/goph/fxt/debug"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
)

// HandlerParams provides a set of dependencies for a prometheus http handler constructor.
type HandlerParams struct {
	dig.In

	Handler  debug.Handler
	Gatherer prometheus.Gatherer

	Logger promhttp.Logger      `optional:"true"`
	Opts   promhttp.HandlerOpts `optional:"true"`
}