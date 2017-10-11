package fxt_test

import (
	"fmt"

	"github.com/goph/fxt"
	"go.uber.org/fx"
)

func ExampleLifecycle() {
	var ctx struct {
		Closer fxt.Closer
	}

	type A struct{}

	fx.New(
		fx.Provide(fxt.NewLifecycle),
		fx.Provide(func(l fxt.Lifecycle) *A {
			l.Append(fxt.Hook{
				// OnStart and OnStop are valid fx hooks as well.
				OnClose: func() error {
					fmt.Print("closing")

					return nil
				},
			})

			return &A{}
		}),
		fx.Extract(&ctx),
		fx.Invoke(func(a *A) {}),
		fx.NopLogger,
	)
	defer ctx.Closer.Close()

	// Output: closing
}
