package fxsql

import (
	"fmt"
	"time"
)

// Config holds a list of options used during the debug server construction.
type Config struct {
	Driver string
	Dsn    string

	// See https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime
	ConnMaxLifetime time.Duration

	// Zero means no change (default), negative means 0, positive means itself
	// See https://golang.org/pkg/database/sql/#DB.SetMaxIdleConns
	MaxIdleConns int

	// See https://golang.org/pkg/database/sql/#DB.SetMaxOpenConns
	MaxOpenConns int
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
