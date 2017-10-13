package grpc

// Config holds a list of options used during the grpc server construction.
type Config struct {
	Network string
	Addr    string

	// Register the reflection API or not
	ReflectionEnabled bool
}

// NewConfig returns a new config populated with default values.
func NewConfig(addr string) *Config {
	return &Config{
		Network: "tcp",
		Addr:    addr,
	}
}
