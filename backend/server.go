package main

import (
	"flag"
	db "socialnetwork/pkg/db/sqlite"
	"socialnetwork/pkg/models"
)

const DATABASE_FILE_PATH = "./pkg/db/"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	if *dbinit {
		dbName := flag.Arg(0)
		dbFilePath := &models.BasicDatabaseInit{
			Directory: DATABASE_FILE_PATH,
			DBName:    dbName,
		}
		db.InitialiseDatabase(dbFilePath)

	}

}
