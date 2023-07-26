package dbutils

import (
	"database/sql"
	"fmt"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/helpermodels"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

/*
OpenDatabase takes a specific file path to a database and opens it.

The function will first check if the faile path to a database is
valid. If it is, it will then open a database and keep that connection alive to be used for
the life of the server.

Parameters:
  - filepath (*helpermodels.FilePathComponents): a pointer to a struct of file path components

Returns:
  - error: An error is returned if the database fails to open

Example:
  - The database will open when dbinit is called in the terminal with the name of the database.
*/
func OpenDatabase(filepath *helpermodels.FilePathComponents) error {
	var err error
	dbFilePath := path.Join(filepath.Directory, filepath.FileName+filepath.Extension)

	isValidPath, err := helpers.IsValidPath(dbFilePath)
	if !isValidPath {
		return fmt.Errorf("path to database not valid: %s", err)
	}

	DB, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to open %s database: %s", filepath.FileName, err)
	}

	return nil
}

/*
CloseDatabase closes an open database.

The function will take an existing open database and close it.

Example:
  - The closing of the open database will be defered until the server closes.
*/
func CloseDatabase() {
	DB.Close()
}
