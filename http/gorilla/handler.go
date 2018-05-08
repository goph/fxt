package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewHandler constructs a new HTTP handler instance.
func NewHandler(router *mux.Router) http.Handler {
	return router
}
