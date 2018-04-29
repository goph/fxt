package correlationid

import "net/http"

//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdSource -output . -outpkg correlationid_test -testonly"

// correlationIdSource retrieves a correlation ID from a request.
type correlationIdSource interface {
	// ExtractCorrelationID extracts the correlation ID from a request.
	// When a correlation ID cannot be found, it returns an empty string.
	ExtractCorrelationID(r *http.Request) string
}
