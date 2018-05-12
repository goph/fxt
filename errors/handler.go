package fxerrors

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	emperror_log "github.com/goph/emperror/log"
	"go.uber.org/dig"
)

// HandlerParams provides a set of dependencies for an error handler constructor.
type HandlerParams struct {
	dig.In

	Logger   log.Logger         `optional:"true"`
	Handlers []emperror.Handler `group:"error"`
}

// HandlerResult can be used to provide a Handler attached to the error group.
type HandlerResult struct {
	dig.Out

	Handler emperror.Handler `group:"error"`
}

// NewHandler returns a new error handler.
func NewHandler(params HandlerParams) emperror.Handler {
	handlers := params.Handlers

	// Configure a log handler if a logger is provided
	if params.Logger != nil {
		handlers = append(handlers, emperror_log.NewHandler(params.Logger))
	}

	var handler emperror.Handler

	// Check if a composite handler is necessary
	if len(handlers) == 0 {
		handler = emperror.NewNopHandler()
	} else if len(handlers) == 1 {
		handler = handlers[0]
	} else {
		handler = emperror.NewCompositeHandler(handlers...)
	}

	return handler
}
