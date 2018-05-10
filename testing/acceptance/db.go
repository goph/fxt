package acceptance

import (
	"database/sql"
)

// DbContext can be used to expose the database from the application.
type DbContext struct {
	DB *sql.DB
}
