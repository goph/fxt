package fxopentracing

import "github.com/opentracing/opentracing-go"

// NewTracer returns the registered global tracer.
func NewTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}
