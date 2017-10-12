package http

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"go.uber.org/dig"
)

// ServerParams provides a set of dependencies for a http server constructor.
type ServerParams struct {
	dig.In

	Config    *Config
	Handler   http.Handler
	Logger    log.Logger `optional:"true"`
	Lifecycle fxt.Lifecycle
}
