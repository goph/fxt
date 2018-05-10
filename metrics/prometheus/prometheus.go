package fxprometheus

import (
	"os"

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

// RegisterDefaultCollectors registers default metric collectors in a prometheus registerer instance.
func RegisterDefaultCollectors(registerer prometheus.Registerer) error {
	err := registerer.Register(prometheus.NewProcessCollector(os.Getpid(), ""))
	if err != nil {
		return err
	}

	err = registerer.Register(prometheus.NewGoCollector())
	if err != nil {
		return err
	}

	return nil
}

// RegisterCollectors registers collector instances in the registerer.
func RegisterCollectors(params CollectorParams) error {
	for _, collector := range params.Collectors {
		err := params.Registerer.Register(collector)
		if err != nil {
			return err
		}
	}

	return nil
}
