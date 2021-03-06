package correlationid_test

import (
	"testing"

	"context"

	fxcontext "github.com/goph/fxt/context"
	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUnaryServerInterceptor(t *testing.T) {
	source := new(correlationIdSource)

	ctx := context.Background()

	chain := grpc_middleware.ChainUnaryServer(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			source.On("ExtractCorrelationID", ctx).Return("cid", true)

			return handler(ctx, req)
		},
		correlationid.UnaryServerInterceptor(correlationid.WithSource(source)),
	)

	var called bool

	chain(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		called = true

		cid, ok := fxcontext.CorrleationId(ctx)

		assert.True(t, ok)
		assert.Equal(t, "cid", cid)

		return nil, nil
	})

	assert.True(t, called)
	source.AssertExpectations(t)
}
