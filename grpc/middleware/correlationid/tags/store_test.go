package tagscorrelationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/goph/fxt/grpc/middleware/correlationid/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestTagStore_StoreCorrelationID(t *testing.T) {
	source := new(correlationIdSource)

	ctx := context.Background()

	chain := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			source.On("ExtractCorrelationID", ctx).Return("cid", true)

			return handler(ctx, req)
		},
		correlationid.UnaryServerInterceptor(
			correlationid.WithSource(source),
			correlationid.WithStore(tagscorrelationid.NewTagStore()),
		),
	)

	var called bool

	chain(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		called = true

		tags := grpc_ctxtags.Extract(ctx)

		assert.Equal(t, "cid", tags.Values()[correlationid.TagCorrelationID])

		return nil, nil
	})

	assert.True(t, called)
	source.AssertExpectations(t)
}
