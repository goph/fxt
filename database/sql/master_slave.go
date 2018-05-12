package fxsql

import (
	"database/sql"

	"github.com/goph/fxt"
	"github.com/goph/healthz"
	"go.uber.org/dig"
)

// MasterSlaveConfigResult can be used to register master and slave configurations in the container.
type MasterSlaveConfigResult struct {
	dig.Out

	Master *Config `name:"master"`
	Slave  *Config `name:"slave"`
}

// MasterSlaveConnectionParams provides a set of dependencies for a master-slave database connection constructor.
type MasterSlaveConnectionParams struct {
	dig.In

	MasterConfig    *Config           `name:"master"`
	SlaveConfig     *Config           `name:"slave"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}

// MasterSlaveConnectionResult contains the result connections of the NewMasterSlaveConnection constructor.
type MasterSlaveConnectionResult struct {
	dig.Out

	Master *sql.DB `name:"master"`
	Slave  *sql.DB `name:"slave"`
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
