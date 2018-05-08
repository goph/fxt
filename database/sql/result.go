package fxsql

import (
	"database/sql"

	"go.uber.org/dig"
)

// MasterSlaveConnectionResult contains the result connections of the NewMasterSlaveConnection constructor.
type MasterSlaveConnectionResult struct {
	dig.Out

	Master *sql.DB `name:"master"`
	Slave  *sql.DB `name:"slave"`
}

// MasterSlaveConfigResult can be used to register master and slave configurations in the container.
type MasterSlaveConfigResult struct {
	dig.Out

	Master *Config `name:"master"`
	Slave  *Config `name:"slave"`
}
