package airbrake

import (
	"net/http"

	"github.com/airbrake/gobrake"
)

// Config holds a list of options used during the airbrake handler construction.
type Config struct {
	ProjectID  int64
	ProjectKey string

	Host    string
	Filters []func(notice *gobrake.Notice) *gobrake.Notice
	Async   bool

	HttpClient *http.Client
}

// NewConfig returns a new config populated with default values.
func NewConfig(projectId int64, projectKey string) *Config {
	return &Config{
		ProjectID:  projectId,
		ProjectKey: projectKey,
	}
}

// AppConfig can be used in an application config to represent Airbrake connection details.
// It supports github.com/goph/nest
type AppConfig struct {
	Enabled    bool   `env:""`
	Endpoint   string `env:""`
	ProjectID  int64  `env:"project_id"`
	ProjectKey string `env:"" split_words:"true"`
}
