package correlationid

import (
	"net/http"

	"github.com/goph/fxt/context"
	"github.com/goph/fxt/internal/correlationid"
)

type middleware struct {
	headers   []string
	generator correlationIdGenerator
}

// New returns a new correlation ID middleware.
func New(options ...Option) *middleware {
	m := new(middleware)

	for _, o := range options {
		o.apply(m)
	}

	if m.generator == nil {
		m.generator = correlationid.NewRandGenerator(32)
	}

	return m
}

func (m *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, ok := context.CorrleationId(ctx)
		if !ok { // Do not overwrite existing correlation ID
			ok := false

			for _, header := range m.headers { // Check headers (if any)
				if cid := r.Header.Get(header); cid != "" {
					ok = true

					ctx = context.WithCorrelationId(r.Context(), cid)
					r = r.WithContext(ctx)
				}
			}

			if !ok { // No header found: generate one
				cid := m.generator.Generate()

				ctx = context.WithCorrelationId(r.Context(), cid)
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(rw, r)
	})
}
