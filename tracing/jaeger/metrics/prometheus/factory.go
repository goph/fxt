package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uber/jaeger-lib/metrics"
	jaegerprometheus "github.com/uber/jaeger-lib/metrics/prometheus"
)

// NewMetricsFactory creates a new prometheus metrics factory for Jaeger.
func NewMetricsFactory(registerer prometheus.Registerer) metrics.Factory {
	return jaegerprometheus.New(jaegerprometheus.WithRegisterer(registerer))
}
