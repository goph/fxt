package test

import (
	"context"

	"github.com/DATA-DOG/godog"
	"github.com/goph/fxt"
)

// AppContext adds scenario hooks to a Godog suite.
func AppContext(app *fxt.App) func(s *godog.Suite) {
	return func(s *godog.Suite) {
		s.BeforeScenario(func(scenario interface{}) {
			err := app.Start(context.Background())
			if err != nil {
				panic(err)
			}
		})

		s.AfterScenario(func(scenario interface{}, err error) {
			err = app.Stop(context.Background())
			if err != nil {
				panic(err)
			}

			app.Close()
		})
	}
}
