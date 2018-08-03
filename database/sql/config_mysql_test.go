package fxsql_test

import (
	"testing"

	"github.com/goph/fxt/database/sql"
	"github.com/stretchr/testify/assert"
)

func TestMysqlAppConfig_Driver(t *testing.T) {
	config := fxsql.MysqlAppConfig{}

	assert.Equal(t, "mysql", config.Driver())
}

func TestMysqlAppConfig_DSN(t *testing.T) {
	config := fxsql.MysqlAppConfig{
		Host: "host",
		Port: 3306,
		User: "root",
		Pass: "",
		Name: "database",
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	dsn := config.DSN()

	assert.Equal(t, "root:@tcp(host:3306)/database?parseTime=true", dsn)
}
