package jaeger

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/client/log/go-kit"
)

// NewTracer returns a new jaeger tracer.
func NewTracer(params TracerParams) (opentracing.Tracer, error) {
	var jaegerOptions []jaegercfg.Option

	if params.Logger != nil {
		jaegerOptions = append(jaegerOptions, jaegercfg.Logger(xkit.NewLogger(log.With(params.Logger, "component", "jaeger"))))
	}

	if params.MetricsFactory != nil {
		jaegerOptions = append(jaegerOptions, jaegercfg.Metrics(params.MetricsFactory))
	}

	tracer, closer, err := params.Config.JaegerConfig.New(params.Config.ServiceName, jaegerOptions...)
	if err != nil {
		return nil, err
	}

	params.Lifecycle.Append(fxt.Hook{
		OnClose: closer.Close,
	})

	return tracer, nil
}
