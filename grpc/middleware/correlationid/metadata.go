package correlationid

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

type metadataSource struct {
	headers []string
}

// NewMetadataSource returns a correlation ID source which retrieves the correlation ID from gRPC metadata headers.
func NewMetadataSource(headers ...string) *metadataSource {
	return &metadataSource{
		headers: headers,
	}
}

// DefaultMetadataSource returns a correlation ID source which retrieves the correlation ID from gRPC metadata headers.
func DefaultMetadataSource() *metadataSource {
	return NewMetadataSource("correlation_id", "correlation-id", "correlationid")
}

func (s *metadataSource) ExtractCorrelationID(ctx context.Context) string {
	md := metautils.ExtractIncoming(ctx)

	for _, header := range s.headers { // Check headers (if any)
		if cid := md.Get(header); cid != "" {
			return cid
		}
	}

	return ""
}
