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

func (s *headerSource) ExtractCorrelationID(r *http.Request) string {
	for _, header := range s.headers { // Check headers (if any)
		if cid := r.Header.Get(header); cid != "" {
			return cid
		}
	}

	return ""
}
