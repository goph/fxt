package debug_test

import (
	"net"

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
			func() *debug.Config {
				addr, _ := net.ResolveIPAddr("tcp", ":8080")
				return debug.NewConfig(addr)
			},
			debug.NewHealthCollector,
			debug.NewServer,
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
