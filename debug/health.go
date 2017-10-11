package debug

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
