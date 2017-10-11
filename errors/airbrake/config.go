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
