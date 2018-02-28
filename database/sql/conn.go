package sql

import (
	"database/sql"

	"github.com/goph/fxt"
	"github.com/goph/healthz"
)

// NewConnection creates a new database connection and optionally registers a health check.
// It also registers a closer in the application lifecycle.
func NewConnection(params ConnectionParams) (*sql.DB, error) {
	db, err := sql.Open(params.Config.Driver, params.Config.Dsn)
	if err != nil {
		return nil, err
	}

	// Connection options
	db.SetConnMaxLifetime(params.Config.ConnMaxLifetime)
	db.SetMaxIdleConns(params.Config.MaxIdleConns)
	db.SetMaxOpenConns(params.Config.MaxOpenConns)

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))
	}

	params.Lifecycle.Append(fxt.Hook{
		OnClose: db.Close,
	})

	return db, err
}

// NewMasterSlaveConnection calls NewConnection twice with different input configurations.
func NewMasterSlaveConnection(params MasterSlaveConnectionParams) (MasterSlaveConnectionResult, error) {
	result := MasterSlaveConnectionResult{}

	db, err := NewConnection(ConnectionParams{
		Config:          params.MasterConfig,
		HealthCollector: params.HealthCollector,
		Lifecycle:       params.Lifecycle,
	})
	if err != nil {
		return result, err
	}
	result.Master = db

	db, err = NewConnection(ConnectionParams{
		Config:          params.SlaveConfig,
		HealthCollector: params.HealthCollector,
		Lifecycle:       params.Lifecycle,
	})
	if err != nil {
		return result, err
	}
	result.Master = db

	return result, err
}
