package fxt

import "go.uber.org/fx"

// App wraps the original fx.App instance and adds a closer method to it.
type App struct {
	*fx.App

	closer Closer
}

// New creates a new App and in turn an fx.App instance.
func New(opts ...fx.Option) *App {
	app := new(App)
	opts = append(opts, Bootstrap, fx.Populate(&app.closer))

	app.App = fx.New(opts...)

	return app
}

// Close invokes the OnClose hook of the extended lifecycle.
func (a *App) Close() error {
	return a.closer.Close()
}

// ApplicationInfo is an optional set of information that can be set by the runtime environment (eg. console application).
type ApplicationInfo struct {
	Version    string
	CommitHash string
	BuildDate  string
}
