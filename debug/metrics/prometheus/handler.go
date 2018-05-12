package fxdbgprometheus

import (
	"github.com/goph/fxt/debug"
	"github.com/goph/fxt/metrics/prometheus/promhttp"
)

// RegisterHandler registers the prometheus HTTP handler in the debug mux.
func RegisterHandler(h fxdebug.Handler, p fxpromhttp.Handler) {
	h.Handle("/metrics", p)
}
