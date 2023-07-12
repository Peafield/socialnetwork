package db

import (
	"database/sql"
	"log"
	"os"
	"path"
)

/*
InitialiseDatabase initialises the database.

It defines a file path to where the database should be stored. It then checks
if the the database already exists at the this file path. If it does not, it then
creates the file. It then opens the database using an sqlite3 driver and sets the foreign keys
to be on. It then closes the database.

Parameters:
  - name: the database name as a string

Errors:
  - if the file path is invalid.
  - if the file fails to be created.
  - if the database fails to open.

Example:
  - InitialiseDatabase is only used once, called when the database is initially created.
*/
func InitialiseDatabase(dir string, name string) {
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
