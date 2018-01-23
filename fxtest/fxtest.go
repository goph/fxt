package fxtest

import (
	"github.com/goph/fxt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// App wraps the original fxtest.App instance.
type App struct {
	*fxtest.App

	closer fxt.Closer
}

// New creates a new test application.
func New(tb fxtest.TB, opts ...fx.Option) *App {
	app := new(App)
	opts = append(opts, fxt.Bootstrap, fx.Populate(&app.closer))

	app.App = fxtest.New(tb, opts...)

	return app
}

// Close invokes the OnClose hook of the extended lifecycle.
func (a *App) Close() error {
	return a.closer.Close()
}
