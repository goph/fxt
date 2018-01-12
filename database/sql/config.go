package sql

import "fmt"

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

// AppConfig can be used in an application config to represent database connection details.
// It supports github.com/goph/nest
// It is recommended to use a prefix:
//		type Config struct {
//			Db sql.AppConfig `prefix:"db"`
//		}
type AppConfig struct {
	Host string `env:"" required:"true"`
	Port int    `env:"" default:"3306"`
	User string `env:"" required:"true"`
	Pass string `env:""` // Required removed for now because empty value is not supported by Viper
	Name string `env:"" required:"true"`
}

// Dsn returns the DSN created form the configuration.
func (c AppConfig) Dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
	)
}
