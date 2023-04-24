package database

import (
	"fmt"
)

func (s *Service) dbConnectionString(dbType, name string) string {
	switch dbType {
	case typeSqlite, "sqlite":
		return s.sqliteConnectionString(name)
	case typeMysql:
		return s.mysqlConnectionString()
	}

	return ""
}

func (s *Service) sqliteConnectionString(name string) string {
	var (
		user  = s.Config().Get(keyDatabaseUser)
		pass  = s.Config().Get(keyDatabasePass)
		crypt = s.Config().Get(keyDatabaseCrypt)
		salt  = s.Config().Get(keyDatabaseSalt)
	)

	fmtStr := "file:%s?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s"
	result := fmt.Sprintf(fmtStr, name, user, pass, crypt)

	switch crypt {
	case cryptSaltedSHA1, cryptSaltedSHA256, cryptSaltedSHA384, cryptSaltedSHA512:
		result += fmt.Sprintf("&_auth_salt=%s", salt)
	}

	return result
}

func (s *Service) mysqlConnectionString() string {
	var (
		protocol = s.Config().Get(keyDatabaseProtocol)
		host     = s.Config().Get(keyDatabaseHost)
		port     = s.Config().Get(keyDatabasePort)
		name     = s.Config().Get(keyDatabaseName)
		user     = s.Config().Get(keyDatabaseUser)
		pass     = s.Config().Get(keyDatabasePass)
	)

	const fmtStr = "%v:%v@%v(%v:%v)/%v?multiStatements=true&parseTime=true&interpolateParams=true"
	return fmt.Sprintf(fmtStr, user, pass, protocol, host, port, name)
}
