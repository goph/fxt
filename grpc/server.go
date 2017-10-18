package grpc

import (
	"context"
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// NewServer creates a new grpc server.
func NewServer(params ServerParams) (*grpc.Server, Err) {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "grpc")

	// TODO: separate log levels
	// TODO: only set logger once
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(
		log.NewStdlibAdapter(level.Info(logger)),
		log.NewStdlibAdapter(level.Warn(logger)),
		log.NewStdlibAdapter(level.Error(logger)),
	))

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewTCPChecker(params.Config.Addr))
	}

	options := params.Config.Options

	if params.StreamInterceptor != nil {
		options = append(options, grpc.StreamInterceptor(params.StreamInterceptor))
	}

	if params.UnaryInterceptor != nil {
		options = append(options, grpc.UnaryInterceptor(params.UnaryInterceptor))
	}

	server := grpc.NewServer(options...)

	if params.Config.ReflectionEnabled {
		level.Debug(logger).Log("msg", "grpc reflection service enabled")

		reflection.Register(server)
	}

	errCh := make(chan error, 1)

	params.Lifecycle.Append(fxt.Hook{
		OnStart: func(ctx context.Context) error {
			level.Info(logger).Log(
				"msg", "listening on address",
				"addr", params.Config.Addr,
				"network", params.Config.Network,
			)

			lis, err := net.Listen(params.Config.Network, params.Config.Addr)
			if err != nil {
				return err
			}

			go func() {
				errCh <- server.Serve(lis)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return fxt.WithTimeout(ctx, func(ctx context.Context) error {
				server.GracefulStop()

				return nil
			})
		},
		OnClose: func() error {
			server.Stop()

			return nil
		},
	})

	return server, errCh
}
