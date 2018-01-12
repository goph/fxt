package sql_test

import (
	"testing"

	"github.com/goph/fxt/database/sql"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig_Dsn(t *testing.T) {
	config := sql.AppConfig{
		Host: "host",
		Port: 3306,
		User: "root",
		Pass: "",
		Name: "database",
	}

	dsn := config.Dsn()

	assert.Equal(t, "root:@tcp(host:3306)/database", dsn)
}
