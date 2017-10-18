package grpc

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"go.uber.org/dig"
)

// ServerParams provides a set of dependencies for a grpc server constructor.
type ServerParams struct {
	dig.In

	Config          *Config
	Logger          log.Logger        `optional:"true"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}
