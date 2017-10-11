package airbrake

import (
	"github.com/airbrake/gobrake"
	"github.com/goph/emperror/airbrake"
)

// NewHandler returns a new Airbrake handler.
func NewHandler(config *Config) *airbrake.Handler {
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

	return &airbrake.Handler{
		Notifier:          notifier,
		SendSynchronously: !config.Async,
	}
}
