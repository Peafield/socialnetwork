package db

import (
	"database/sql"
	"log"
	"os"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

/*
InitialiseDatabase validates the name and the file path of the given database.

If the conditions are met, proceed to create the database.

Parameters:
  - dbFilePath: An interface containing the basic requirements to initialise a database.

Errors:
  - If the db name is not valid, exit the program and log the error.
  - If the db directory file path is not valid, exit the program and log the error.
*/
func InitialiseDatabase(dbFilePath models.DatabaseInit) {
	dbName := dbFilePath.GetDBName()
	dbDirectory := dbFilePath.GetDirectory()

	isDBNameValid, err := helpers.IsAlphaNumeric(dbName)

	if !isDBNameValid {
		log.Fatalf("DB name contains non alpha-numeric characters. Err: %s", err)
	}

	isFilePathValid, err := helpers.IsValidPath(dbDirectory)

	if !isFilePathValid {
		log.Fatalf("DB directory is not valid. Err: %s", err)
	}

	CreateDatabase(dbDirectory, dbName)

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
func CreateDatabase(dir string, name string) {
	filepath := path.Join(dir, name+".db")
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			log.Fatalf("failed to create file path: %s", err)
		}
		file.Close()
	} else {
		log.Printf("File path error: %s", err)
	}
	db, err := sql.Open("sqlite3", filepath+"/?_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()
}
