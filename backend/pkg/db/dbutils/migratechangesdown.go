package db

import (
	"fmt"
	"log"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

/*
MigrateChangesDown apply the current down migrations to the database.

It will first validate that the filepath is correct and then apply any
current up migrations to the database. If no changes have taken place, or an error
occurs and error will be returned.

Parameters:
  - dbFilePath (DatabaseManager): contains methods to retrieve directory and file name of a database.

Returns:
  - error: if the file path is not valid; migration initialisaing failed; migration failed.
*/
func MigrateChangesDown(dbFilePath dbmodels.DatabaseManager) error {
	dbDir := dbFilePath.GetDirectory()
	dbName := dbFilePath.GetDBName() + ".db"
	filePath := path.Join(dbDir, dbName)

	isValidPath, err := helpers.IsValidPath(filePath)
	if !isValidPath {
		return fmt.Errorf("file path is not valid. Err: %s", err)
	}

	m, err := migrate.New(
		"file://./pkg/db/migrations",
		"sqlite3://"+filePath)
	if err != nil {
		return fmt.Errorf("migration initialization failed: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("down migration failed: %v", err)
	} else {
		log.Println("Migration succeeded")
	}
	return nil
}
