package prometheus

import (
	"github.com/goph/fxt/debug"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterHandler registers the Prometheus metrics handler in the debug server.
func RegisterHandler(handler debug.Handler) {
	handler.Handle("/metrics", promhttp.Handler())
}
