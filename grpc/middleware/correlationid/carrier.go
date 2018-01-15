package correlationid

import "golang.org/x/net/context"

type Carrier interface {
	// GetCorrelationID returns an existing correlation ID.
	// If the second argument is false, the interceptor attempts to set the correlation ID.
	// If the correlation ID is empty, the second argument should be false in order to generate a new one.
	GetCorrelationID(ctx context.Context) (string, bool)

	// SetCorrelationID sets a correlation ID in the context and returns the new context (if any).
	SetCorrelationID(ctx context.Context, correlationID string) context.Context
}
