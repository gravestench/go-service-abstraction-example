package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

const (
	typeSqlite = "sqlite"
	typeMysql  = "mysql"

	cryptSaltedSHA1   = "SSHA1"
	cryptSaltedSHA256 = "SSHA256"
	cryptSaltedSHA384 = "SSHA384"
	cryptSaltedSHA512 = "SSHA512"

	keyDatabaseType     = "type"
	defaultDatabaseType = "sqlite3"

	keyDatabaseName     = "dbname"
	defaultDatabaseName = ""

	keyDatabaseUser     = "user"
	defaultDatabaseUser = ""

	keyDatabasePass     = "pass"
	defaultDatabasePass = ""

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

type Service struct {
	log abstract.Logger
	cfg abstract.ValueControl

	db *sqlx.DB
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.populateDependencies(possibleDependencies)
	s.setupConfig()
	s.setupDatabase()
}

func (s *Service) Name() string {
	return "Database"
}

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
		s.cfg.Default(k, v)
	}
}

func (s *Service) setupDatabase() {
	dbType := s.cfg.Get(keyDatabaseType)

	db, err := sqlx.Open(dbType, s.dbConnectionString(dbType))
	if err == nil {
		s.db = db
	}

	switch dbType {
	case typeSqlite:

	case typeMysql:
	}

	s.log.Fatal().Msgf("opening database: %v", err)
}

func (s *Service) dbConnectionString(dbType string) string {
	switch dbType {
	case typeSqlite:
		return s.sqliteConnectionString()
	case typeMysql:
		return s.mysqlConnectionString()
	}

	return ""
}

func (s *Service) sqliteConnectionString() string {
	var (
		name  = s.cfg.Get(keyDatabaseName)
		user  = s.cfg.Get(keyDatabaseUser)
		pass  = s.cfg.Get(keyDatabasePass)
		crypt = s.cfg.Get(keyDatabaseCrypt)
		salt  = s.cfg.Get(keyDatabaseSalt)
	)

	fmtStr := "file:%s.s3db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s"
	result := fmt.Sprintf(fmtStr, name, user, pass, crypt)

	switch crypt {
	case cryptSaltedSHA1, cryptSaltedSHA256, cryptSaltedSHA384, cryptSaltedSHA512:
		result += fmt.Sprintf("&_auth_salt=%s", salt)
	}

	return result
}

func (s *Service) mysqlConnectionString() string {
	var (
		protocol = s.cfg.Get(keyDatabaseProtocol)
		host     = s.cfg.Get(keyDatabaseHost)
		port     = s.cfg.Get(keyDatabasePort)
		name     = s.cfg.Get(keyDatabaseName)
		user     = s.cfg.Get(keyDatabaseUser)
		pass     = s.cfg.Get(keyDatabasePass)
	)

	const fmtStr = "%v:%v@%v(%v:%v)/%v?multiStatements=true&parseTime=true&interpolateParams=true"
	return fmt.Sprintf(fmtStr, user, pass, protocol, host, port, name)
}
