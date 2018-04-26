package opentracing

import (
	"net/http"

	"github.com/goph/fxt/context"
	"github.com/opentracing/opentracing-go"
)

type middleware struct {
	restrictedKey string
}

// New returns a new middleware that sets a correlation ID (if any) in an opentracing span (if any).
func New(options ...Option) *middleware {
	m := new(middleware)

	for _, o := range options {
		o.apply(m)
	}

	// Default restricted key
	if m.restrictedKey == "" {
		m.restrictedKey = defaultRestrictedKey
	}

	return m
}

func (m *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			cid := span.BaggageItem(m.restrictedKey)
			if cid == "" { // Do not overwrite existing correlation ID
				if cid, ok := context.CorrleationId(ctx); ok {
					span := opentracing.SpanFromContext(ctx)
					if span != nil {
						span.SetBaggageItem(m.restrictedKey, cid)
					}
				}
			}
		}

		next.ServeHTTP(rw, r)
	})
}
