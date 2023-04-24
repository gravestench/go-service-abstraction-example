package database

import (
	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

//go:embed schema
var fs embed.FS

type Service struct {
	log abstract.Logger
	cfg abstract.ConfigurationManager

	flags struct {
		sqliteFilePath string
	}

	db *sqlx.DB
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Init(app abstract.ServiceManager) {
	s.setupConfig()
	s.setupDatabase()
	s.printDebugInfo()
}

func (s *Service) Name() string {
	return "Database"
}

func (s *Service) createMysqlDatabase() {
	s.log.Fatal().Msgf("[%s] mysql not yet implemented", s.Name())
}
