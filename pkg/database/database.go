package database

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
)

func (s *Service) Database() *sqlx.DB {
	return s.db
}

func (s *Service) setupDatabase() {
	var (
		dbType     = s.Config().Get(keyDatabaseType)
		dbFileName = s.Config().Get(keyDatabaseName)
	)

	// override with the supplied cli flag if it was specified
	if s.flags.sqliteFilePath != "" {
		dbFileName = s.flags.sqliteFilePath
	}

	dbDir := s.cfg.DirPath()
	dbPath := filepath.Join(dbDir, dbFileName)

	if !fileExists(dbPath) {
		switch dbType {
		case typeSqlite, "sqlite":
			s.createSqliteDatabase(dbPath)
		case typeMysql:
			s.createMysqlDatabase()
		}
	}

	if db, err := sqlx.Open(dbType, s.dbConnectionString(dbType, dbPath)); err == nil {
		s.log.Info().Msgf("connected to '%s' database '%s'", dbType, dbPath)
		s.db = db
		return
	}

	if db, err := sqlx.Open(dbType, s.dbConnectionString(dbType, dbPath)); err != nil {
		s.log.Fatal().Msgf("could not connect to databse: %v", err)
		return
	} else {
		s.db = db
	}
}

func (s *Service) printDebugInfo() {
	type count struct {
		Count string `db:"count"`
	}

	tables := []string{
		"user_role",
		"user",
		"logs",
		"alarms",
		"ueattachlogs",
		"radio",
		"widget",
		"system_settings",
		"terms",
		"nssai_profile",
		"smf_selection_profile",
		"sm_policy_data_profile",
		"connected_users",
		"utilization",
	}

	time.Sleep(time.Second)

	for _, table := range tables {
		res := make([]count, 0)

		query := fmt.Sprintf("select count(*) as count from %s", table)

		if err := s.db.Select(&res, query, table); err != nil {
			s.log.Fatal().Msgf("getting %s info: %v", table, err)
		}

		s.log.Debug().Msgf("entries in '%s' table: %v", table, res[0].Count)
	}
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
