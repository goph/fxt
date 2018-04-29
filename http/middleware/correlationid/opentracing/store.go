package otcorrelationid

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
)

type opentracingStore struct {
	restrictedKey string
}

// NewOpentracingStore returns a correlation ID source and store using opentracing trace.
func NewOpentracingStore(options ...Option) *opentracingStore {
	s := new(opentracingStore)
	op := newOptions(options...)

	s.restrictedKey = op.restrictedKey

	return s
}

func (p *opentracingStore) StoreCorrelationID(r *http.Request, cid string) {
	ctx := r.Context()

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetBaggageItem(p.restrictedKey, cid)
	}
}
