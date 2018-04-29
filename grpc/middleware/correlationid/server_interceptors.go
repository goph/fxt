package correlationid

import (
	"context"

	fxcontext "github.com/goph/fxt/context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

const TagCorrelationID = "correlationid"

// UnaryServerInterceptor returns a new unary server interceptor for propagating correlation ID.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := newOptions(opts...)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = serverCorrelationID(o, ctx)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for propagating correlation ID.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := newOptions(opts...)

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := serverCorrelationID(o, stream.Context())

		wrappedStream := grpc_middleware.WrapServerStream(stream)
		wrappedStream.WrappedContext = ctx

		err := handler(srv, wrappedStream)

		return err
	}
}

func serverCorrelationID(opts *options, ctx context.Context) context.Context {
	cid, ok := fxcontext.CorrleationId(ctx)
	if ok { // Do not overwrite existing correlation ID
		for _, store := range opts.stores {
			store.StoreCorrelationID(ctx, cid)
		}

		return ctx
	}

	for _, source := range opts.sources {
		cid = source.ExtractCorrelationID(ctx)
		if cid != "" {
			break
		}
	}

	if cid != "" { // A correlation ID was found
		for _, store := range opts.stores {
			store.StoreCorrelationID(ctx, cid)
		}

		ctx = fxcontext.WithCorrelationId(ctx, cid)
	}

	return ctx
}
