package http

import (
	"context"
	stdlog "log"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
)

// NewServer creates a new http server.
func NewServer(params ServerParams) Err {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "http")

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
				return err
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
