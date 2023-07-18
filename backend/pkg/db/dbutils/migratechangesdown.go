package dbutils

import (
	"fmt"
	"log"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"

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
  - dbFilePath (FilePathManager): contains methods to retrieve directory, extension, and file name of a database.
  - migrationsPath (string): a relative path to the migrations folder.
  - migrationConstructor (MigrationConstructor): an interface to abstract new migrations.
  - migrateUpDown (MigrationUpdates): an interface to abstract migration database changes up or down.

Returns:
  - error: if the file path is not valid; migration initialisaing failed; migration failed.
*/
func MigrateChangesDown(dbFilePath helpermodels.FilePathManager, migrationsPath string, migrationConstructor dbmodels.MigrationConstructor, migrateUpDown dbmodels.MigrationUpdates) error {
	dbDir := dbFilePath.GetDirectory()
	dbName := dbFilePath.GetFileName()
	dbExtension := dbFilePath.GetFileExtension()
	filePath := path.Join(dbDir, dbName+dbExtension)

	isValidPath, err := helpers.IsValidPath(filePath)
	if !isValidPath {
		return fmt.Errorf("file path is not valid. Err: %s", err)
	}

	m, err := migrationConstructor.New(
		"file://"+migrationsPath,
		"sqlite3://"+filePath)
	if err != nil {
		return fmt.Errorf("migration initialization failed: %v, %v", err, filePath)
	}

	if err := migrateUpDown.Down(m); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("down migration failed: %v", err)
	} else {
		log.Println("Migration succeeded")
	}
	return nil
}
