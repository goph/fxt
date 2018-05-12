package fxsql

import (
	"database/sql"

	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"go.uber.org/dig"
)

// ConnectionParams provides a set of dependencies for a database connection constructor.
type ConnectionParams struct {
	dig.In

	Config          *Config
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}

// NewConnection creates a new database connection and optionally registers a health check.
// It also registers a closer in the application lifecycle.
func NewConnection(params ConnectionParams) (*sql.DB, error) {
	db, err := sql.Open(params.Config.Driver, params.Config.Dsn)
	if err != nil {
		return nil, err
	}

	// Connection options
	db.SetConnMaxLifetime(params.Config.ConnMaxLifetime)
	db.SetMaxOpenConns(params.Config.MaxOpenConns)

	// Setting the default value of the field (which is zero) means no idle connections.
	// To maintain the internal default behavior DB, zero means no change, negative values mean zero (no idle),
	// positive values mean themselves.
	if params.Config.MaxIdleConns != 0 {
		db.SetMaxIdleConns(params.Config.MaxIdleConns)
	}

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))
	}

	params.Lifecycle.Append(fxt.Hook{
		OnClose: db.Close,
	})

	return db, err
}
