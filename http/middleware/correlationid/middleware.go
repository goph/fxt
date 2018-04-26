package correlationid

import (
	"net/http"

	"github.com/goph/fxt/context"
	"github.com/goph/fxt/internal/correlationid"
)

const (
	defaultHeader = "Correlation-ID"
)

// Option sets an option in the middleware.
type Option func(*middleware)

// Header sets the header in the middleware.
func Header(header string) Option {
	return func(m *middleware) {
		m.header = header
	}
}

type middleware struct {
	generator correlationid.Generator

	header string
}

// New returns a new middleware.
func New(generator correlationid.Generator, options ...Option) *middleware {
	m := &middleware{
		generator: generator,
	}

	for _, o := range options {
		o(m)
	}

	// Default header
	if m.header == "" {
		m.header = defaultHeader
	}

	return m
}

func (m *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get(m.header)
		if cid == "" {
			cid = m.generator.Generate()
		}

		ctx := context.WithCorrelationId(r.Context(), cid)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
