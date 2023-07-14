package main

import (
	"flag"
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
)

const DATABASE_FILE_PATH = "./pkg/db/"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	if *dbinit {
		dbName := flag.Arg(0)
		dbFilePath := &dbmodels.DatabaseFilePathComponents{
			Directory: DATABASE_FILE_PATH,
			DBName:    dbName,
		}
		db.InitialiseDatabase(dbFilePath)
	}

	dbup := flag.Bool("dbup", false, "Migrate database changes up")
	if *dbup {
		dbName := flag.Arg(0)
		dbFilePath := &dbmodels.DatabaseFilePathComponents{
			Directory: DATABASE_FILE_PATH,
			DBName:    dbName,
		}
		db.MigrateChangesUp(dbFilePath)
	}
}
