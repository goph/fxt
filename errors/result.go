package fxerrors

import (
	"github.com/goph/emperror"
	"go.uber.org/dig"
)

// HandlerResult can be used to provide a Handler attached to the error group.
type HandlerResult struct {
	dig.Out

	Handler emperror.Handler `group:"error"`
}
