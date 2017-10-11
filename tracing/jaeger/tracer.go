package jaeger

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer returns a new opentracing tracer.
func NewTracer(params TracerParams) (opentracing.Tracer, error) {
	var jaegerOptions []jaegercfg.Option

	if params.Logger != nil {
		jaegerOptions = append(jaegerOptions, jaegercfg.Logger(&kitLogger{params.Logger}))
	}

	// TODO: handle closer
	tracer, closer, err := params.Config.JaegerConfig.New(params.Config.ServiceName, jaegerOptions...)
	if err != nil {
		return nil, err
	}

	params.Lifecycle.Append(fxt.Hook{
		OnClose: closer.Close,
	})

	return tracer, nil
}

// kitLogger wraps the application logger instance in a Jaeger compatible one.
type kitLogger struct {
	logger log.Logger
}

// Error implements the github.com/uber/jaeger-client-go/log.Logger interface.
func (l *kitLogger) Error(msg string) {
	level.Error(l.logger).Log("msg", msg)
}

// Infof implements the github.com/uber/jaeger-client-go/log.Logger interface.
func (l *kitLogger) Infof(msg string, args ...interface{}) {
	level.Info(l.logger).Log("msg", fmt.Sprintf(msg, args...))
}
