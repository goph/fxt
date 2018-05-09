package fxhttp_test

import (
	"net/http"

	"github.com/goph/fxt"
	fxhttp "github.com/goph/fxt/http"
	"go.uber.org/fx"
)

func Example() {
	app := fx.New(
		fx.NopLogger,
		fxt.Bootstrap,
		fx.Provide(
			func() *fxhttp.Config {
				return fxhttp.NewConfig(":8080")
			},
			func() http.Handler {
				return http.NewServeMux()
			},
			fxhttp.NewServer,
		),
	)

	if err := app.Err(); err != nil {
		panic(err)
	}

	// Output:
}
