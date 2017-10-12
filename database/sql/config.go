package sql

// Config holds a list of options used during the debug server construction.
type Config struct {
	Driver string
	Dsn    string
}

// NewConfig returns a new config populated with default values.
func NewConfig(driver string, dsn string) *Config {
	return &Config{
		Driver: driver,
		Dsn:    dsn,
	}
}
