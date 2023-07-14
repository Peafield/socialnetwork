package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	Helpers models.Helper
}

/*
InitialiseDatabase validates the name and the file path of the given database.

If the conditions are met, proceed to create the database.

Parameters:
  - dbFilePath: An interface containing the basic requirements to initialise a database.

Returns:
  - error: returns a specified error

Errors:
  - If the db name is not valid, exit the program and log the error.
  - If the db directory file path is not valid, exit the program and log the error.
*/
func InitialiseDatabase(dbFilePath models.DatabaseInit) error {
	dbName := dbFilePath.GetDBName()
	dbDirectory := dbFilePath.GetDirectory()

	isDBNameValid, err := helpers.IsAlphaNumeric(dbName)

	if !isDBNameValid {
		return errors.New(fmt.Sprintf("DB name contains non alpha-numeric characters. Err: %s", err))
	}

	isFilePathValid, err := helpers.IsValidPath(dbDirectory)

	if !isFilePathValid {
		return errors.New(fmt.Sprintf("DB directory is not valid. Err: %s", err))
	}

	return CreateDatabase(dbDirectory, dbName)
}

/*
CreateDatabase initialises the database.

It defines a file path to where the database should be stored. It then checks
if the the database already exists at the this file path. If it does not, it then
creates the file. It then opens the database using an sqlite3 driver and sets the foreign keys
to be on. It then closes the database.

Parameters:
  - dir: the directory file path in which the database should be created
  - name: the database name as a string

Errors:
  - if the file path is invalid.
  - if the file fails to be created.
  - if the database fails to open.

Example:
  - CreateDatabase is only used once, called when the database is initially created.
*/
func CreateDatabase(dir string, name string) error {
	filepath := path.Join(dir, name+".db")

	//removed file path validator as we do it in the helpers

	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		return errors.New(fmt.Sprintf("failed to create file path: %s", err))
	}

	db, err := sql.Open("sqlite3", filepath+"/?_foreign_keys=on")
	defer db.Close()

	if err != nil {
		return errors.New(fmt.Sprintf("failed to open database: %s", err))
	}

	return nil
}
