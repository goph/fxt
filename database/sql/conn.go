package fxsql

import (
	"database/sql"

	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"go.uber.org/dig"
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
	result.Slave = db

	return result, err
}

// MakeMasterPrimaryConnection makes the master connection the primary one to be used as *sql.DB.
func MakeMasterPrimaryConnection(result struct {
	dig.In

	Master *sql.DB `name:"master"`
}) (*sql.DB) {
	return result.Master
}
