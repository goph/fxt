package correlationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/goph/fxt/internal/correlationid/mocks"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUnaryServerInterceptor(t *testing.T) {
	generator := new(mocks.Generator)
	carrier := new(mocks.Carrier)

	ctx := context.Background()

	chain := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			carrier.On("GetCorrelationID", ctx).Return("1234", true)

			return handler(ctx, req)
		},
		correlationid.UnaryServerInterceptor(generator, carrier),
	)

	var called bool

	chain(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		called = true

		tags := grpc_ctxtags.Extract(ctx)

		assert.Equal(t, "1234", tags.Values()[correlationid.TagCorrelationID])

		return nil, nil
	})

	assert.True(t, called)
	generator.AssertNotCalled(t, "Generate")
	carrier.AssertExpectations(t)
}

func TestUnaryServerInterceptor_Empty(t *testing.T) {
	generator := new(mocks.Generator)
	carrier := new(mocks.Carrier)

	generator.On("Generate").Return("1234")

	ctx := context.Background()

	chain := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			carrier.On("GetCorrelationID", ctx).Return("", false)
			carrier.On("SetCorrelationID", ctx, "1234").Return(ctx)

			return handler(ctx, req)
		},
		correlationid.UnaryServerInterceptor(generator, carrier),
	)

	var called bool

	chain(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		called = true

		tags := grpc_ctxtags.Extract(ctx)

		assert.Equal(t, "1234", tags.Values()[correlationid.TagCorrelationID])

		return nil, nil
	})

	assert.True(t, called)
	generator.AssertExpectations(t)
	carrier.AssertExpectations(t)
}
