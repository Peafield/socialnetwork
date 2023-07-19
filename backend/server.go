package main

import (
	"flag"
	"log"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
)

const DATABASE_FILE_PATH = "./pkg/db/"
const MIGRATIONS_FILE_PATH = "./pkg/db/migrations"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	dbup := flag.Bool("dbup", false, "Migrate database changes up")
	dbdown := flag.Bool("dbdown", false, "Migrate database changes down")

	flag.Parse()

	if *dbinit {
		dbName := flag.Arg(0)
		if len(dbName) < 1 {
			log.Fatalf("Missing database name")
		}
		dbFilePath := &helpermodels.FilePathComponents{
			Directory: DATABASE_FILE_PATH,
			FileName:  dbName,
			Extension: ".db",
		}
		err := dbutils.CreateDatabase(dbFilePath)
		if err != nil {
			log.Fatalf("Failed to initialise database: %s", err)
		}

	}

	if *dbup {
		dbName := flag.Arg(0)
		dbFilePath := &helpermodels.FilePathComponents{
			Directory: DATABASE_FILE_PATH,
			FileName:  dbName,
			Extension: ".db",
		}
		migrationConstructor := &dbmodels.NativeMigrate{}
		migrateUpDown := &dbmodels.NativeMigrateUpdates{}
		err := dbutils.MigrateChangesUp(dbFilePath, MIGRATIONS_FILE_PATH, migrationConstructor, migrateUpDown)
		if err != nil {
			log.Fatalf("Failed to migrate changes up: %s", err)
		}
	}

	if *dbdown {
		dbName := flag.Arg(0)
		dbFilePath := &helpermodels.FilePathComponents{
			Directory: DATABASE_FILE_PATH,
			FileName:  dbName,
			Extension: ".db",
		}
		migrationConstructor := &dbmodels.NativeMigrate{}
		migrateUpDown := &dbmodels.NativeMigrateUpdates{}
		err := dbutils.MigrateChangesDown(dbFilePath, MIGRATIONS_FILE_PATH, migrationConstructor, migrateUpDown)
		if err != nil {
			log.Fatalf("Failed to migrate changes down: %s", err)
		}
	}
	// db, _ := sql.Open("sqlite3", "./pkg/db/socialNetwork.db")
	// user := &dbmodels.User{
	// 	UserId:         "1",
	// 	IsLoggedIn:     0,
	// 	Email:          "user@test.com",
	// 	HashedPassword: "hashed_password",
	// 	FirstName:      "First",
	// 	LastName:       "Last",
	// 	DOB:            time.Now(),
	// 	AvatarPath:     "path/to/avatar",
	// 	DisplayName:    "User",
	// 	AboutMe:        "About me",
	// }
	// userDB.InsertUser(db, user)
}
