package correlationid

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const TagCorrelationID = "correlationid"

// UnaryServerInterceptor returns a new unary server interceptor for propagating correlation ID.
func UnaryServerInterceptor(idGenerator IdGenerator, carrier Carrier) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = serverCorrelationID(idGenerator, carrier, ctx)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for propagating correlation ID.
func StreamServerInterceptor(idGenerator IdGenerator, carrier Carrier) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := serverCorrelationID(idGenerator, carrier, stream.Context())

		wrappedStream := grpc_middleware.WrapServerStream(stream)
		wrappedStream.WrappedContext = ctx

		err := handler(srv, wrappedStream)

		return err
	}
}

func serverCorrelationID(idGenerator IdGenerator, carrier Carrier, ctx context.Context) context.Context {
	correlationID, ok := carrier.GetCorrelationID(ctx)
	if !ok {
		if correlationID == "" {
			correlationID = idGenerator.Generate()
		}

		ctx = carrier.SetCorrelationID(ctx, correlationID)
	}

	// Use tags as the source of correlation ID in handlers.
	tags := grpc_ctxtags.Extract(ctx)
	tags.Set(TagCorrelationID, correlationID)

	return ctx
}
