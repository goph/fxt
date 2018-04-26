package correlationid

import (
	"net/http"

	"github.com/goph/fxt/context"
)

type middleware struct {
	headers []string
}

// New returns a new correlation ID middleware.
func New(options ...Option) *middleware {
	m := &middleware{
		headers: []string{defaultHeader},
	}

	for _, o := range options {
		o.apply(m)
	}

	return m
}

func (m *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		for _, header := range m.headers {
			if cid := r.Header.Get(header); cid != "" {
				ctx := context.WithCorrelationId(r.Context(), cid)
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(rw, r)
	})
}
