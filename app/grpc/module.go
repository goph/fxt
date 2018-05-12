package fxgrpcapp

import (
	"github.com/goph/fxt/app/default"
	"github.com/goph/fxt/grpc"
	"go.uber.org/fx"
)

// Module is an fx compatible module.
var Module = fx.Options(
	fxdefaultapp.Module,

	fx.Provide(fxgrpc.NewServer),
)
