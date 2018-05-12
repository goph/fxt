package fxprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// New registers new instances of registerer and gatherer in the container.
func New() (prometheus.Registerer, prometheus.Gatherer) {
	registry := prometheus.NewRegistry()

	return registry, registry
}

// NewWithGlobal registers new instances of registerer and gatherer in the container.
//
// It also merges the gatherer instance with the global one.
func NewWithGlobal() (prometheus.Registerer, prometheus.Gatherer) {
	registerer, gatherer := New()
	gatherer = prometheus.Gatherers{gatherer, prometheus.DefaultGatherer}

	return registerer, gatherer
}
