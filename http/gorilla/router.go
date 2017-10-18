package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

// Router embeds mux.Router and forces middlewares to be injected for all routes.
type Router struct {
	*mux.Router
	tracer opentracing.Tracer
}

// NewRouter returns a new router instance
func NewRouter(tracer opentracing.Tracer) *Router {
	return &Router{mux.NewRouter(), tracer}
}

// Handle registers a new route with a matcher for the URL path.
func (r *Router) Handle(path string, handler http.Handler) *mux.Route {
	return r.NewRoute().Path(path).Handler(r.traceHandler(handler))
}

// HandleFunc registers a new route with a matcher for the URL path.
func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.NewRoute().Path(path).Handler(r.traceHandler(http.HandlerFunc(f)))
}

func (r *Router) traceHandler(handler http.Handler) http.Handler {
	opName := nethttp.OperationNameFunc(func(r *http.Request) string {
		route := mux.CurrentRoute(r)

		name := route.GetName()

		if name == "" {
			name, _ = route.GetPathTemplate()
		}

		return name
	})

	return nethttp.Middleware(r.tracer, handler, opName)
}
