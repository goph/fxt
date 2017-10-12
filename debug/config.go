package debug

// Config holds a list of options used during the debug server construction.
type Config struct {
	Debug   bool
	Network string
	Addr    string
}

// NewConfig returns a new config populated with default values.
func NewConfig(addr string) *Config {
	return &Config{
		Network: "tcp",
		Addr:    addr,
	}
}
