package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

// opName extracts the operation name from a request.
func opName(r *http.Request) string {
	route := mux.CurrentRoute(r)

	name := route.GetName()

	if name == "" {
		name, _ = route.GetPathTemplate()
	}

	return name
}

// InjectTracer injects the tracer into every route of the router.
//
// Make sure to invoke it after all routes are registered.
func InjectTracer(router *mux.Router, tracer opentracing.Tracer) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		handler := route.GetHandler()
		handler = nethttp.Middleware(tracer, handler, nethttp.OperationNameFunc(opName))
		route.Handler(handler)

		return nil
	})
}
