package errors

import (
	"github.com/go-kit/kit/log"
	"go.uber.org/dig"
)

// HandlerParams provides a set of dependencies for an error handler constructor.
type HandlerParams struct {
	dig.In

	Logger   log.Logger   `optional:"true"`
	Handlers HandlerStack `optional:"true"`
}
