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
MigrateChangesUp apply the current up migrations to the database.

It will first validate that the filepath is correct and then apply any
current up migrations to the database. If no changes have taken place, or an error
occurs and error will be returned.

Parameters:
  - dbFilePath (DatabaseManager): contains methods to retrieve directory and file name of a database.
  - migrationConstructor (MirgationConstructor): contains methods to initialize a new migration
  - migrateUpdDown (MigrationUpdates): contains methods to update a specific migration (receiver type is a pointer)

Returns:
  - error: if the file path is not valid; migration initialisaing failed; migration failed.
*/
func MigrateChangesUp(dbFilePath dbmodels.DatabaseManager, migrationConstructor dbmodels.MigrationConstructor, migrateUpDown dbmodels.MigrationUpdates) error {
	dbDir := dbFilePath.GetDirectory()
	dbName := dbFilePath.GetDBName() + ".db"
	filePath := path.Join(dbDir, dbName)

	isValidPath, err := helpers.IsValidPath(filePath)
	if !isValidPath {
		return fmt.Errorf("file path is not valid. Err: %s", err)
	}

	m, err := migrationConstructor.New(
		"file://./pkg/db/migrations",
		"sqlite3://"+filePath)
	if err != nil {
		return fmt.Errorf("migration initialization failed: %v, %v", err, filePath)
	}

	if err := migrateUpDown.Up(m); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("up migration failed: %v", err)
	} else {
		log.Println("Migration succeeded")
	}
	return nil
}
