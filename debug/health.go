package fxdebug

import (
	"github.com/goph/healthz"
)

// NewHealthCollector returns a new healthz.Collector.
func NewHealthCollector(handler Handler) healthz.Collector {
	healthCollector := healthz.Collector{}

	// Add health checks
	handler.Handle("/healthz", healthCollector.Handler(healthz.LivenessCheck))
	handler.Handle("/readiness", healthCollector.Handler(healthz.ReadinessCheck))

	return healthCollector
}

// NewStatusChecker returns a new healthz.StatusChecker with "Healthy" as the default value.
// It also registers the checker in the health collector.
func NewStatusChecker(collector healthz.Collector) *healthz.StatusChecker {
	status := healthz.NewStatusChecker(healthz.Healthy)

	collector.RegisterChecker(healthz.ReadinessCheck, status)

	return status
}
