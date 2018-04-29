package correlationid

import "context"

//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdSource -output . -outpkg correlationid_test -testonly"
//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdSource -output ./tags -outpkg tagscorrelationid_test -testonly"

// correlationIdSource retrieves a correlation ID from a request.
type correlationIdSource interface {
	// ExtractCorrelationID extracts the correlation ID from a context.
	// When a correlation ID cannot be found, it returns an empty string.
	ExtractCorrelationID(ctx context.Context) string
}
