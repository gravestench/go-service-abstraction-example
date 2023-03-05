package abstract

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseService interface {
	Database() *sqlx.DB
}
