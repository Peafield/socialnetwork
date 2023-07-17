package db

import (
	"fmt"
	"log"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"

	_ "github.com/mattn/go-sqlite3"
)

/*
InitialiseDatabase validates the name and the file path of the given database.

If the conditions are met, proceed to create the database.

Parameters:
  - dbFilePath (string): An interface containing the basic requirements to initialise a database.

Returns:
  - error: returns a specified error

Errors:
  - If the db name is not valid, exit the program and log the error.
  - If the db directory file path is not valid, exit the program and log the error.
  - If the db already exists, exit the program and log the error.
*/
func InitialiseDatabase(dbFilePath dbmodels.DatabaseManager) error {
	dbName := dbFilePath.GetDBName()
	dbDirectory := dbFilePath.GetDirectory()

	isDBNameValid, err := helpers.IsAlphaNumeric(dbName)
	log.Println(isDBNameValid)
	if !isDBNameValid {
		return fmt.Errorf("DB name contains non alpha-numeric characters. Err: %s", err)
	}

	isValidDirPath, err := helpers.IsValidPath(dbDirectory)

	if !isValidDirPath {
		return fmt.Errorf("database directory is not valid. Err: %s", err)
	}

	fullPath := path.Join(dbDirectory, dbName+".db")
	dbExists, _ := helpers.IsValidPath(fullPath)

	if dbExists {
		return fmt.Errorf("database already exists")
	}

	osFileCreator := helpermodels.OSFileCreator{}

	SQLDBOpener := dbmodels.SQLDBOpener{
		DriveName:      "sqlite3",
		DataSourceName: dbDirectory + "/?_foreign_keys=on",
	}

	return CreateDatabase(dbDirectory, dbName, &osFileCreator, &SQLDBOpener)
}

/*
CreateDatabase initialises the database.

It defines a file path to where the database should be stored. It then checks
if the the database already exists at the this file path. If it does not, it then
creates the file. It then opens the database using an sqlite3 driver and sets the foreign keys
to be on. It then closes the database.

Parameters:
  - dir (string): the directory file path in which the database should be created
  - name (string): the database name
  - fileCreator (struct):  has a Create() method
  - dbOpener (struct): has a Open() method

Errors:
  - if the file path is invalid.
  - if the file fails to be created.
  - if the database fails to open.

Example:
  - CreateDatabase is only used once, called when the database is initially created.
*/

func CreateDatabase(dir string, name string, fileCreator helpermodels.FileCreator, dbOpener dbmodels.DBOpener) error {
	filepath := path.Join(dir, name+".db")

	file, err := fileCreator.Create(filepath)
	log.Println(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file path: %s", err)
	}
	defer file.Close()

	db, err := dbOpener.Open("sqlite3", filepath+"/?_foreign_keys=on")
	if err != nil {
		return fmt.Errorf("failed to open database: %s", err)
	}
	defer db.Close()

	return nil
}
