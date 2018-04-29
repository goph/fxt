package otcorrelationid

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type opentracingSource struct {
	restrictedKey string
}

// NewOpentracingSource returns a correlation ID source using opentracing trace.
func NewOpentracingSource(options ...Option) *opentracingSource {
	s := new(opentracingSource)
	op := newOptions(options...)

	s.restrictedKey = op.restrictedKey

	return s
}

func (s *opentracingSource) ExtractCorrelationID(ctx context.Context) string {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		return span.BaggageItem(s.restrictedKey)
	}

	return ""
}
