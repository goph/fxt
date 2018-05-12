package fxjaeger

// AppConfig can be used in an application config to represent Jaeger connection details.
// It supports github.com/goph/nest
type AppConfig struct {
	Enabled bool   `env:""`
	Addr    string `env:""`
}
