package debug

import (
	"context"
	"expvar"
	stdlog "log"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt"
	"golang.org/x/net/trace"
)

// NewServer creates a new debug server.
func NewServer(params ServerParams) (Handler, Err) {
	handler := http.NewServeMux()

	if params.Config.Debug {
		handler.Handle("/debug/vars", expvar.Handler())
		handler.HandleFunc("/debug/pprof/", pprof.Index)
		handler.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		handler.HandleFunc("/debug/pprof/profile", pprof.Profile)
		handler.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		handler.HandleFunc("/debug/pprof/trace", pprof.Trace)
		handler.HandleFunc("/debug/requests", trace.Traces)
		handler.HandleFunc("/debug/events", trace.Events)

		// This is probably okay, as this service should not be exposed to public in the first place.
		trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
			return true, true
		}
	}

	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "debug")

	server := &http.Server{
		Handler:  handler,
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

	return handler, errCh
}
