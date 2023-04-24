package abstract

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DatabaseService describes a service which contains the app's database
// session, and exposes some helper methods for exec'ing database queries.
// most services will likely use the `Database()` method to check for
// a non-nil database instance to prevent themselves from init'ing while it
// is not ready.
type DatabaseService interface {
	Database() *sqlx.DB
	Get(result interface{}, query string, args ...interface{}) error
	Select(result interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}
