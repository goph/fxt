package fxdebug_test

import (
	"github.com/goph/fxt"
	"github.com/goph/fxt/debug"
	"github.com/goph/healthz"
	"go.uber.org/fx"
)

func Example() {
	status := healthz.NewStatusChecker(healthz.Healthy)

	app := fx.New(
		fx.NopLogger,
		fxt.Bootstrap,
		fx.Provide(
			func() *fxdebug.Config {
				return fxdebug.NewConfig(":8080")
			},
			fxdebug.NewHealthCollector,
			fxdebug.NewServer,
		),
		fx.Invoke(func(collector healthz.Collector) {
			collector.RegisterChecker(healthz.ReadinessCheck, status)
		}),
	)

	if err := app.Err(); err != nil {
		panic(err)
	}

	// Output:
}
