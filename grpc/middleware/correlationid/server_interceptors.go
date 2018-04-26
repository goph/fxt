package correlationid

import (
	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid/internal"
	"github.com/goph/fxt/internal/correlationid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

const TagCorrelationID = "correlationid"

// UnaryServerInterceptor returns a new unary server interceptor for propagating correlation ID.
func UnaryServerInterceptor(generator correlationid.Generator, carrier internal.Carrier) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = serverCorrelationID(generator, carrier, ctx)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for propagating correlation ID.
func StreamServerInterceptor(generator correlationid.Generator, carrier internal.Carrier) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := serverCorrelationID(generator, carrier, stream.Context())

		wrappedStream := grpc_middleware.WrapServerStream(stream)
		wrappedStream.WrappedContext = ctx

		err := handler(srv, wrappedStream)

		return err
	}
}

func serverCorrelationID(generator correlationid.Generator, carrier internal.Carrier, ctx context.Context) context.Context {
	correlationID, ok := carrier.GetCorrelationID(ctx)
	if !ok {
		if correlationID == "" {
			correlationID = generator.Generate()
		}

		ctx = carrier.SetCorrelationID(ctx, correlationID)
	}

	// Use tags as the source of correlation ID in handlers.
	tags := grpc_ctxtags.Extract(ctx)
	tags.Set(TagCorrelationID, correlationID)

	return ctx
}
