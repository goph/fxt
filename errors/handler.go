package errors

import (
	"github.com/goph/emperror"
	emperror_log "github.com/goph/emperror/log"
)

// NewHandler returns a new error handler.
func NewHandler(params HandlerParams) emperror.Handler {
	var handlers []emperror.Handler

	if params.Logger != nil {
		handlers = append(handlers, emperror_log.NewHandler(params.Logger))
	}

	var handler emperror.Handler

	if len(handlers) == 0 {
		handler = emperror.NewNopHandler()
	} else if len(handlers) == 1 {
		handler = handlers[0]
	} else {
		handler = emperror.NewCompositeHandler(handlers...)
	}

	return handler
}
