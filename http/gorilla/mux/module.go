package fxmux

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// Module is an fx compatible module.
var Module = fx.Provide(mux.NewRouter, NewHandler)
