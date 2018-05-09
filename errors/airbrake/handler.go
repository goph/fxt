package fxairbrake

import (
	"github.com/airbrake/gobrake"
	"github.com/goph/emperror/airbrake"
	"github.com/goph/fxt"
)

// NewHandler returns a new Airbrake handler.
func NewHandler(config *Config, lifecycle fxt.Lifecycle) *airbrake.Handler {
	notifier := gobrake.NewNotifier(config.ProjectID, config.ProjectKey)

	if config.Host != "" {
		notifier.SetHost(config.Host)
	}

	if config.HttpClient != nil {
		notifier.Client = config.HttpClient
	}

	for _, filter := range config.Filters {
		notifier.AddFilter(filter)
	}

	handler := &airbrake.Handler{
		Notifier:          notifier,
		SendSynchronously: !config.Async,
	}

	lifecycle.Append(fxt.Hook{
		OnClose: handler.Close,
	})

	return handler
}
