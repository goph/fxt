package fxdefaultapp

import (
	"github.com/goph/fxt/debug"
	"github.com/goph/fxt/errors"
	"github.com/goph/fxt/log"
	"go.uber.org/fx"
)

// Module is an fx compatible module.
var Module = fx.Options(
	fx.Provide(
		// Log and error handling
		fxlog.NewLogger,
		fxerrors.NewHandler,

		// Debug server
		fxdebug.NewServer,
		fxdebug.NewHealthCollector,
		fxdebug.NewStatusChecker,
	),
)
