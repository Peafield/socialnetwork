package dbutils

import (
	"fmt"
	"log"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/helpermodels"

	_ "github.com/mattn/go-sqlite3"
)

/*
CreateDatabase initializes the creation of a SQLite database file.

It takes as input an object that implements the FilePathManager interface. This interface
should provide methods to retrieve the directory and file name of the database file.
It constructs a full path from these components, checks if a valid file can be created at
that location, and if so, creates the file.

Parameters:
  - dbFilePath (helpermodels.FilePathManager): An object that provides the directory and
    file name to construct the full path for the SQLite database file.

Errors:
  - Returns an error if the file path is invalid or if the database file fails to be created.

Example:
  - CreateDatabase is called when the application starts, to ensure that a valid SQLite
    database file is available for further database operations.
*/
func CreateDatabase(dbFilePath helpermodels.FilePathManager) error {
	isValidDatabaseFilePath, err := helpers.CheckValidPath(dbFilePath)
	if !isValidDatabaseFilePath {
		return fmt.Errorf("invalid database: %s", err)
	}

	dbFullPath := path.Join(dbFilePath.GetDirectory(), dbFilePath.GetFileName()+dbFilePath.GetFileExtension())
	log.Println("Create database:", dbFullPath)
	osFileCreator := &helpermodels.OSFileCreator{}

	err = helpers.FileCreator(dbFullPath, osFileCreator)
	if err != nil {
		return fmt.Errorf("failed to create database file: %s", err)
	}

	return nil
}
