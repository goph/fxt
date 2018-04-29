package tagscorrelationid

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

type tagStore struct {
	tag string
}

// NewTagStore returns a correlation ID store using gRPC tags.
func NewTagStore(options ...Option) *tagStore {
	s := new(tagStore)

	for _, o := range options {
		o.apply(s)
	}

	// Default tag
	if s.tag == "" {
		s.tag = TagCorrelationID
	}

	return s
}

func (s *tagStore) StoreCorrelationID(ctx context.Context, cid string) {
	tags := grpc_ctxtags.Extract(ctx)

	tags.Set(s.tag, cid)
}
