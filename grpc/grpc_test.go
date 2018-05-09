package fxgrpc_test

import (
	"github.com/goph/fxt"
	fxgrpc "github.com/goph/fxt/grpc"
	"go.uber.org/fx"
)

func Example() {
	app := fx.New(
		fx.NopLogger,
		fxt.Bootstrap,
		fx.Provide(
			func() *fxgrpc.Config {
				return fxgrpc.NewConfig(":8080")
			},
			fxgrpc.NewServer,
		),
	)

	if err := app.Err(); err != nil {
		panic(err)
	}

	// Output:
}
