package otmux

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

type TracingParams struct {
	fx.In

	Router  *mux.Router
	Tracer  opentracing.Tracer
	Options []nethttp.MWOption `group:"default"`
}

// InjectTracer injects the tracer into every route of the router.
//
// Make sure to invoke it after all routes are registered.
func InjectTracer(params TracingParams) {
	options := append([]nethttp.MWOption{nethttp.OperationNameFunc(opName)}, params.Options...)

	params.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		handler := route.GetHandler()

		handler = nethttp.Middleware(params.Tracer, handler, options...)

		route.Handler(handler)

		return nil
	})
}

// opName extracts the operation name from a request.
func opName(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route == nil {
		return r.Proto + " " + r.Method
	}

	if name := route.GetName(); name != "" {
		return name
	}

	if tpl, err := route.GetPathTemplate(); err == nil {
		return r.Proto + " " + r.Method + " " + tpl
	}

	return r.Proto + " " + r.Method
}
