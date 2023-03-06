package database

import (
	"flag"
)

const (
	flagSqlitePath            = "sqlite-path"
	flagSqlitePathDefault     = "db.s3db"
	flagSqlitePathDescription = "path to sqlite database"
)

func (s *Service) Flags() []string {
	return []string{
		flagSqlitePath,
	}
}

func (s *Service) Parse(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&s.flags.sqliteFilePath, flagSqlitePath, flagSqlitePathDefault, flagSqlitePathDescription)

	return fs.Parse(args)
}
