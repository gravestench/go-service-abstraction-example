package abstract

import (
	"github.com/jmoiron/sqlx"
)

// DatabaseService is not very well implemented, i wanted it to be
// generic regarding the actual backing database.
type DatabaseService interface {
	Database() *sqlx.DB
}
