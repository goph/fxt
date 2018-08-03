package fxsql

import (
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

// connDetails describes a structure that contains everything necessary for creating a database connection.
type connDetails interface {
	// Driver returns the driver name for the connection.
	Driver() string

	// DSN returns the data source name for the connection.
	DSN() string
}

// NewConfig returns a new config populated with default values.
func NewConfigFromConnectionDetails(details connDetails) *Config {
	return &Config{
		Driver: details.Driver(),
		Dsn:    details.DSN(),
	}
}
