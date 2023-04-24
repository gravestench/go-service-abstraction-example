package database

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

const (
	typeSqlite = "sqlite3"
	typeMysql  = "mysql"

	cryptSaltedSHA1   = "SSHA1"
	cryptSaltedSHA256 = "SSHA256"
	cryptSaltedSHA384 = "SSHA384"
	cryptSaltedSHA512 = "SSHA512"

	keyDatabaseType     = "type"
	defaultDatabaseType = typeSqlite

	keyDatabaseName     = "dbname"
	defaultDatabaseName = "db.s3db"

	keyDatabaseUser     = "user"
	defaultDatabaseUser = "xmc"

	keyDatabasePass     = "pass"
	defaultDatabasePass = "p4ssw0rd"

	keyDatabaseCrypt     = "crypt"
	defaultDatabaseCrypt = ""

	keyDatabaseSalt     = "sqlite.salt"
	defaultDatabaseSalt = ""

	keyDatabaseProtocol     = "mysql.protocol"
	defaultDatabaseProtocol = ""

	keyDatabaseHost     = "mysql.host"
	defaultDatabaseHost = ""

	keyDatabasePort     = "mysql.port"
	defaultDatabasePort = ""
)

func (s *Service) setupConfig() {
	for k, v := range map[string]string{
		keyDatabaseType:     defaultDatabaseType,
		keyDatabaseName:     defaultDatabaseName,
		keyDatabaseUser:     defaultDatabaseUser,
		keyDatabasePass:     defaultDatabasePass,
		keyDatabaseCrypt:    defaultDatabaseCrypt,
		keyDatabaseSalt:     defaultDatabaseSalt,
		keyDatabaseProtocol: defaultDatabaseProtocol,
		keyDatabaseHost:     defaultDatabaseHost,
		keyDatabasePort:     defaultDatabasePort,
	} {
		s.Config().Default(k, v)
	}
}

func (s *Service) Config() abstract.ConfigGroup {
	return s.cfg.GetConfigGroup(s.Name())
}
