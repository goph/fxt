package correlationid

import (
	"net/http"

	"github.com/goph/fxt/context"
)

type middleware struct {
	sources []correlationIdSource
	stores  []correlationIdStore
}

// New returns a new correlation ID middleware.
func New(options ...Option) *middleware {
	m := new(middleware)

	for _, o := range options {
		o.apply(m)
	}

	return m
}

// Retrieves a correlation ID from various sources:
//		- if it is already in the context do not overwrite it, but write it in opentracing (if available)
func (m *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cid, ok := context.CorrleationId(ctx)
		if ok { // Do not overwrite existing correlation ID
			for _, store := range m.stores {
				store.StoreCorrelationID(r, cid)
			}

			next.ServeHTTP(rw, r)
			return
		}

		for _, source := range m.sources {
			cid = source.ExtractCorrelationID(r)
			if cid != "" {
				break
			}
		}

		if cid != "" { // A correlation ID was found
			for _, store := range m.stores {
				store.StoreCorrelationID(r, cid)
			}

			ctx = context.WithCorrelationId(ctx, cid)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(rw, r)
	})
}
