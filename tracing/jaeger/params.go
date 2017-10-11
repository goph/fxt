package jaeger

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"go.uber.org/dig"
)

// TracerParams provides a set of dependencies for a tracer constructor.
type TracerParams struct {
	dig.In

	Config    *Config `optional:"true"`
	Logger    log.Logger
	Lifecycle fxt.Lifecycle
}
