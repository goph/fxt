package fxprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/dig"
)

// CollectorResult can be used to provide a Collector attached to the metrics group.
type CollectorResult struct {
	dig.Out

	Collector prometheus.Collector `group:"default"`
}
