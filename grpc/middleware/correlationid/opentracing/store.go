package otcorrelationid

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type opentracingStore struct {
	restrictedKey string
}

// NewOpentracingStore returns a correlation ID store using opentracing trace.
func NewOpentracingStore(options ...Option) *opentracingStore {
	s := new(opentracingStore)
	op := newOptions(options...)

	s.restrictedKey = op.restrictedKey

	return s
}

func (s *opentracingStore) StoreCorrelationID(ctx context.Context, cid string) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetBaggageItem(s.restrictedKey, cid)
	}
}
