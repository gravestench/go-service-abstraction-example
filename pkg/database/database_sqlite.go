package database

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	"github.com/jmoiron/sqlx"
)

func (s *Service) createSqliteDatabase(dbPath string) {
	if _, err := exec.LookPath("sqlite3"); err != nil {
		s.log.Fatal().Msgf("[%s] sqlite3 not found on PATH ", s.Name())
	}

	s.log.Info().Msgf("[%s] creating default sqlite database", s.Name())

	// Create the database file
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		s.log.Fatal().Msgf("[%s] creating database file: %v", s.Name(), err)
	}

	// Read the schema file
	schema, err := fs.ReadFile("schema/sqlite.sql")
	if err != nil {
		s.log.Fatal().Msgf("[%s] reading embedded(!) schema file: %v", s.Name(), err)
	}

	// Create the tables and indexes defined in the schema file
	_, err = db.Exec(string(schema))
	if err != nil {
		s.log.Fatal().Msgf("[%s] parsing schema file: %s", s.Name(), err)
	}

	s.log.Info().Msgf("[%s] created default sqlite database file '%s'", s.Name(), dbPath)

	db.Close()
}

func (s *Service) validateSchema() {
	// Connect to the SQLite database
	db, err := sqlx.Connect("sqlite3", s.flags.sqliteFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//// Read the schema file
	//schema, err := ioutil.ReadFile("path/to/schema.sql")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//// Check if the database matches the schema
	//tx, err := db.Beginx()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//_, err = tx.Exec(string(schema))
	//if err != nil {
	//	fmt.Println("Schema check failed:", err)
	//	os.Exit(1)
	//}

	//err = tx.Commit()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

}
