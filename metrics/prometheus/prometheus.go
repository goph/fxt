package prometheus

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

// New registers new instances of registerer and gatherer in the container.
func New() (prometheus.Registerer, prometheus.Gatherer, error) {
	registry := prometheus.NewRegistry()

	err := registerDefaultCollectors(registry)
	if err != nil {
		return nil, nil, err
	}

	return registry, registry, nil
}

// registerDefaultCollectors registers default metric collectors in a prometheus registerer instance.
func registerDefaultCollectors(registerer prometheus.Registerer) error {
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
