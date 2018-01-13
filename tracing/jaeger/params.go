package jaeger

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/dig"
)

// TracerParams provides a set of dependencies for a tracer constructor.
type TracerParams struct {
	dig.In

	Config         *Config
	Logger         log.Logger      `optional:"true"`
	MetricsFactory metrics.Factory `optional:"true"`
	Lifecycle      fxt.Lifecycle
}
