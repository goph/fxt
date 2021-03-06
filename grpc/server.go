package fxgrpc

import (
	"context"
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"github.com/goph/fxt/log"
	"github.com/goph/healthz"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// ServerParams provides a set of dependencies for a grpc server constructor.
type ServerParams struct {
	dig.In

	Config          *Config
	Logger          log.Logger        `optional:"true"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}

// NewServer creates a new gRPC server.
func NewServer(params ServerParams) (*grpc.Server, Err) {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "grpc")

	// TODO: separate log levels
	// TODO: only set logger once
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(
		fxlog.NewWriterAdapter(level.Info(logger)),
		fxlog.NewWriterAdapter(level.Warn(logger)),
		fxlog.NewWriterAdapter(level.Error(logger)),
	))

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewTCPChecker(params.Config.Addr))
	}

	server := grpc.NewServer(params.Config.Options...)

	if params.Config.ReflectionEnabled {
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

			if params.Config.ReflectionEnabled {
				level.Debug(logger).Log("msg", "grpc reflection service enabled")
			}

			lis, err := net.Listen(params.Config.Network, params.Config.Addr)
			if err != nil {
				return errors.WithStack(err)
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
