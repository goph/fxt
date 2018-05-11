package fxpromhttp

import "go.uber.org/fx"

// Module is an fx compatible module.
var Module = fx.Provide(
	NewHandler,
	NewLogger,
)
