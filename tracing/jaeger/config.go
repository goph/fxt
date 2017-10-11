package jaeger

import (
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// Config holds a list of options used during the airbrake handler construction.
type Config struct {
	ServiceName string

	JaegerConfig jaegercfg.Configuration
}

// NewConfig returns a new config populated with default values.
func NewConfig(serviceName string) *Config {
	return &Config{
		ServiceName: serviceName,
	}
}
