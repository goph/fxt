package errors

import "github.com/goph/emperror"

// HandlerStack can be used to collect multiple handlers in a single dependecy
// which can be injected into the main error handler constructor.
type HandlerStack []emperror.Handler
