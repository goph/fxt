package fxcontext

import "context"

var contextKeyCorrelationId = contextKey("correlationid")

// WithCorrelationId returns a new child context with added correlation id value.
func WithCorrelationId(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, contextKeyCorrelationId, cid)
}

// CorrleationId gets the correlation id from the context.
func CorrleationId(ctx context.Context) (string, bool) {
	cid, ok := ctx.Value(contextKeyCorrelationId).(string)

	return cid, ok
}
