package errors_test

import (
	"github.com/goph/emperror/airbrake"
	"github.com/goph/fxt/errors"
	airbrakefx "github.com/goph/fxt/errors/airbrake"
	"go.uber.org/dig"
)

func ExampleHandlerStack() {
	dic := dig.New()

	// Provide airbrake config.
	dic.Provide(func() *airbrakefx.Config {
		config := airbrakefx.NewConfig(1, "key")

		return config
	})

	// Provide the airbrake handler itself
	dic.Provide(airbrakefx.NewHandler)

	// Collect all handlers and provide the HandlerStack
	dic.Provide(func(ah *airbrake.Handler) errors.HandlerStack {
		return errors.HandlerStack{ah}
	})

	// The main error constructor will fetch the handler stack
	dic.Provide(errors.NewHandler)
}
