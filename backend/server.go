package main

import (
	"flag"
	"log"
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
)

const DATABASE_FILE_PATH = "./pkg/db/"
const MIGRATIONS_FILE_PATH = "./pkg/db/migrations"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	dbup := flag.Bool("dbup", false, "Migrate database changes up")
	flag.Parse()
	if *dbinit {
		dbName := flag.Arg(0)
		if len(dbName) < 1 {
			log.Fatalf("Missing database name")
		}
		dbFilePath := &dbmodels.DatabaseFilePathComponents{
			Directory: DATABASE_FILE_PATH,
			DBName:    dbName,
		}
		err := db.InitialiseDatabase(dbFilePath)
		if err != nil {
			log.Fatalf("Failed to initialise database: %s", err)
		}

	}

	if *dbup {
		dbName := flag.Arg(0)
		dbFilePath := &dbmodels.DatabaseFilePathComponents{
			Directory: DATABASE_FILE_PATH,
			DBName:    dbName,
		}
		migrationConstructor := &dbmodels.NativeMigrate{}
		migrateUpDown := &dbmodels.NativeMigrateUpdates{}
		err := db.MigrateChangesUp(dbFilePath, MIGRATIONS_FILE_PATH, migrationConstructor, migrateUpDown)
		if err != nil {
			log.Fatalf("Failed to migrate changes up: %s", err)
		}
	}
}
