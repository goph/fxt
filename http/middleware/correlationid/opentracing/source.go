package otcorrelationid

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
)

type opentracingSource struct {
	restrictedKey string
}

// NewOpentracingSource returns a correlation ID source and store using opentracing trace.
func NewOpentracingSource(options ...Option) *opentracingSource {
	s := new(opentracingSource)
	op := newOptions(options...)

	s.restrictedKey = op.restrictedKey

	return s
}

func (p *opentracingSource) ExtractCorrelationID(r *http.Request) string {
	ctx := r.Context()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		return span.BaggageItem(p.restrictedKey)
	}

	return ""
}
