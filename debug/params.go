package debug

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/fxt"
	"go.uber.org/dig"
)

// ServerParams provides a set of dependencies for a debug server constructor.
type ServerParams struct {
	dig.In

	Config    *Config
	Logger    log.Logger `optional:"true"`
	Lifecycle fxt.Lifecycle
}
