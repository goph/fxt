package fxhttp

import (
	"context"
	stdlog "log"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

// ServerParams provides a set of dependencies for a http server constructor.
type ServerParams struct {
	dig.In

	Config          *Config
	Handler         http.Handler
	Logger          log.Logger        `optional:"true"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}

// NewServer creates a new http server.
func NewServer(params ServerParams) Err {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "http")

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewTCPChecker(params.Config.Addr))
	}

	server := &http.Server{
		Handler:  params.Handler,
		ErrorLog: stdlog.New(log.NewStdlibAdapter(level.Error(logger)), "", 0),
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
				return errors.WithStack(err)
			}

			go func() {
				errCh <- server.Serve(lis)
			}()

			return nil
		},
		OnStop:  server.Shutdown,
		OnClose: server.Close,
	})

	return errCh
}
