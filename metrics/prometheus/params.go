package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
)

// CollectorParams collects registered prometheus collectors.
type CollectorParams struct {
	dig.In

	Registerer prometheus.Registerer
	Collectors []prometheus.Collector `group:"default"`
}
