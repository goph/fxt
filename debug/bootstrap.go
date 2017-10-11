package debug

import (
	"go.uber.org/fx"
)

// Bootstrap can be used as an alias for must have provider options.
var Bootstrap = fx.Options(
	fx.Provide(NewHandler),
	fx.Invoke(NewServer),
)
