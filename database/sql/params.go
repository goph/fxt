package fxsql

import (
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

// MasterSlaveConnectionParams provides a set of dependencies for a master-slave database connection constructor.
type MasterSlaveConnectionParams struct {
	dig.In

	MasterConfig    *Config           `name:"master"`
	SlaveConfig     *Config           `name:"slave"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}
