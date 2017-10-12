package debug

import (
	"context"
	stdlog "log"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"github.com/goph/serverz"
	"github.com/goph/stdlib/expvar"
	"github.com/goph/stdlib/net/http/pprof"
	"github.com/goph/stdlib/x/net/trace"
)

// NewServer creates a new debug server.
func NewServer(params ServerParams) (Handler, Err) {
	handler := http.NewServeMux()

	if params.Config.Debug {
		// This is probably okay, as this service should not be exposed to public in the first place.
		trace.SetAuth(trace.NoAuth)

		expvar.RegisterRoutes(handler)
		pprof.RegisterRoutes(handler)
		trace.RegisterRoutes(handler)
	}

	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	server := &serverz.AppServer{
		Server: &http.Server{
			Handler:  handler,
			ErrorLog: stdlog.New(log.NewStdlibAdapter(level.Error(log.With(logger, "server", "debug"))), "", 0),
		},
		Name:   "debug",
		Addr:   params.Config.Addr,
		Logger: logger,
	}

	errCh := make(chan<- error, 1)

	params.Lifecycle.Append(fxt.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen(params.Config.Addr.Network(), params.Config.Addr.String())
			if err != nil {
				return err
			}

			go func() {
				errCh <- server.Serve(lis)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
		OnClose: server.Close,
	})

	return handler, errCh
}
