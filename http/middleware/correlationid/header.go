package correlationid

import "net/http"

type headerSource struct {
	headers []string
}

// NewHeaderSource returns a correlation ID source which retrieves the correlation ID from given HTTP headers.
func NewHeaderSource(headers ...string) *headerSource {
	return &headerSource{
		headers: headers,
	}
}

// DefaultHeaderSource returns a correlation ID source which retrieves the correlation ID from common HTTP headers.
func DefaultHeaderSource() *headerSource {
	return NewHeaderSource("Correlation-ID", "CorrelationID", "X-Correlation-ID")
}

func (s *headerSource) ExtractCorrelationID(r *http.Request) string {
	for _, header := range s.headers { // Check headers (if any)
		if cid := r.Header.Get(header); cid != "" {
			return cid
		}
	}

	return ""
}
