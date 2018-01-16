package correlationid

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"golang.org/x/net/context"
)

const defaultMetadataHeader = "correlationid"

type MetadataCarrierOption func(*metadataCarrier)

// WithHeader customizes the metadata header for the correlation ID.
func WithHeader(header string) MetadataCarrierOption {
	return func(c *metadataCarrier) {
		c.header = header
	}
}

type metadataCarrier struct {
	header string
}

// NewMetadataCarrier creates a carrier operating on gRPC metadata headers.
func NewMetadataCarrier(opts ...MetadataCarrierOption) *metadataCarrier {
	c := new(metadataCarrier)

	for _, opt := range opts {
		opt(c)
	}

	// Default metadata header
	if c.header == "" {
		c.header = defaultMetadataHeader
	}

	return c
}

func (c *metadataCarrier) GetCorrelationID(ctx context.Context) (string, bool) {
	md := metautils.ExtractIncoming(ctx)

	correlationID := md.Get(c.header)

	return correlationID, correlationID != ""
}

func (c *metadataCarrier) SetCorrelationID(ctx context.Context, correlationID string) context.Context {
	md := metautils.ExtractIncoming(ctx).Clone()

	md.Set(c.header, correlationID)

	return md.ToIncoming(ctx)
}

type metadataSourceCarrier struct {
	*metadataCarrier

	carrier Carrier
}

// NewMetadataSourceCarrier returns a carrier which reads the correlation ID from metadata in case the provided carrier
// does not have any configured. Setting the correlation ID on the other hand only sets the underlying carrier,
// but does not set the metadata header.
func NewMetadataSourceCarrier(carrier Carrier, opts ...MetadataCarrierOption) *metadataSourceCarrier {
	return &metadataSourceCarrier{
		metadataCarrier: NewMetadataCarrier(opts...),
		carrier:         carrier,
	}
}

func (c *metadataSourceCarrier) GetCorrelationID(ctx context.Context) (string, bool) {
	correlationID, ok := c.carrier.GetCorrelationID(ctx)
	if !ok {
		correlationID, _ := c.metadataCarrier.GetCorrelationID(ctx)

		return correlationID, false
	}

	return correlationID, true
}

func (c *metadataSourceCarrier) SetCorrelationID(ctx context.Context, correlationID string) context.Context {
	return c.carrier.SetCorrelationID(ctx, correlationID)
}
