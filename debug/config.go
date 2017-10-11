package debug

import (
	"net"
)

// Config holds a list of options used during the debug server construction.
type Config struct {
	Debug bool
	Addr  net.Addr
}

// NewConfig returns a new config populated with default values.
func NewConfig(addr net.Addr) *Config {
	return &Config{
		Addr: addr,
	}
}
