package fxprometheus

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
)

// CollectorParams collects registered prometheus collectors.
type CollectorParams struct {
	dig.In

	Registerer prometheus.Registerer
	Collectors []prometheus.Collector `group:"default"`
}

// CollectorResult can be used to provide a Collector attached to the metrics group.
type CollectorResult struct {
	dig.Out

	Collector prometheus.Collector `group:"default"`
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
