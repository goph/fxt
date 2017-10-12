package fxt

import (
	"go.uber.org/fx"
)

// Bootstrap can be used as an alias for must have provider options composed into one fx option.
var Bootstrap = fx.Provide(NewLifecycle, NewErr)
