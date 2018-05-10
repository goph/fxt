package fxgrpc

import "google.golang.org/grpc"

// Config holds a list of options used during the grpc server construction.
type Config struct {
	Network string
	Addr    string

	// Register the reflection API or not
	ReflectionEnabled bool

	// A list of arbitrary server options for the gRPC server
	Options []grpc.ServerOption
}

// NewConfig returns a new config populated with default values.
func NewConfig(addr string) *Config {
	return &Config{
		Network: "tcp",
		Addr:    addr,
	}
}

// AppClientConfig can be used in an application config to represent gRPC client connection details.
// It supports github.com/goph/nest
type AppClientConfig struct {
	// Addr is the gRPC connection address.
	Addr string `env:"" required:"true"`

	// Host can be used when the connection address and the certificate hostname differs.
	Host string `env:""`

	// Insecure makes the client use insecure channel.
	Insecure bool `env:""`
}
