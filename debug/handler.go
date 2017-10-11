package debug

import (
	"net/http"
)

// Handler accepts an http.Handler or http.HandlerFunc and registers it for a pattern.
type Handler interface {
	http.Handler

	// Handle registers the handler for the given pattern.
	// If a handler already exists for pattern, Handle panics.
	Handle(pattern string, handler http.Handler)

	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// NewHandler creates a new HTTP debug handler.
func NewHandler() Handler {
	return http.NewServeMux()
}
