package opentracing

import (
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
)

const defaultRestrictedKey = "correlation_id"

type Option func(*carrier)

// WithRestrictedKey customizes the baggage item restricted key for the correlation ID.
func WithRestrictedKey(restrictedKey string) Option {
	return func(c *carrier) {
		c.restrictedKey = restrictedKey
	}
}

type carrier struct {
	restrictedKey string
}

// New returns a new opentracing carrier.
func New(opts ...Option) *carrier {
	c := new(carrier)

	for _, opt := range opts {
		opt(c)
	}

	// Default restricted key
	if c.restrictedKey == "" {
		c.restrictedKey = defaultRestrictedKey
	}

	return c
}

func (c *carrier) GetCorrelationID(ctx context.Context) (string, bool) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		grpclog.Info("grpc_correlationid: opentracing span not found for getting correlation id")

		return "", false
	}

	correlationID := span.BaggageItem(c.restrictedKey)

	return correlationID, correlationID == ""
}

func (c *carrier) SetCorrelationID(ctx context.Context, correlationID string) context.Context {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		grpclog.Info("grpc_correlationid: opentracing span not found for setting correlation id")

		return ctx
	}

	span.SetBaggageItem(c.restrictedKey, correlationID)

	return ctx
}
