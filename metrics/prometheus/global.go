package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Global registers the global instances registerer and gatherer in the container.
func Global() (prometheus.Registerer, prometheus.Gatherer) {
	return prometheus.DefaultRegisterer, prometheus.DefaultGatherer
}

// MakeGlobal registers the global instances registerer and gatherer in the container.
//
// Note: this modifies global state which is generally not a good idea.
// Make sure you know what you do when using this option.
func MakeGlobal(registerer prometheus.Registerer, gatherer prometheus.Gatherer) {
	prometheus.DefaultRegisterer = registerer
	prometheus.DefaultGatherer = gatherer
}
