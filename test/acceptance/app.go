package acceptance

import (
	"context"

	"github.com/DATA-DOG/godog"
	"github.com/goph/fxt"
	"go.uber.org/fx"
)

// AppContext restarts the application before every scenario.
type AppContext struct {
	options []fx.Option

	app *fxt.App
}

// NewAppContext returns a new AppContext.
func NewAppContext(options ...fx.Option) *AppContext {
	return &AppContext{
		options: options,
	}
}

// FeatureContext combines Before and After FeatureContext.
func (c *AppContext) FeatureContext(s *godog.Suite) {
	c.BeforeFeatureContext(s)
	c.AfterFeatureContext(s)
}

// BeforeFeatureContext can be called as the first FeatureContext to register application startup as first.
func (c *AppContext) BeforeFeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		app := fxt.New(c.options...)

		if err := app.Err(); err != nil {
			panic(err)
		}

		err := app.Start(context.Background())
		if err != nil {
			panic(err)
		}

		c.app = app
	})
}

// AfterFeatureContext can be called as the last FeatureContext to register application shutdown as last.
func (c *AppContext) AfterFeatureContext(s *godog.Suite) {
	s.AfterScenario(func(interface{}, error) {
		if c.app == nil {
			panic("app not found")
		}

		err := c.app.Stop(context.Background())
		if err != nil {
			panic(err)
		}

		c.app.Close()
	})
}
