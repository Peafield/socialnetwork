package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
)

const DATABASE_FILE_PATH = "./pkg/db/"
const MIGRATIONS_FILE_PATH = "./pkg/db/migrations"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	dbopen := flag.Bool("dbopen", false, "Opens a database and prepares database statements")
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
			log.Fatalf("Failed to create database: %s", err)
		}
	}

	if *dbopen {
		dbName := flag.Arg(0)
		if len(dbName) < 1 {
			log.Fatalf("Missing database name")
		}

		dbFilePath := &helpermodels.FilePathComponents{
			Directory: DATABASE_FILE_PATH,
			FileName:  dbName,
			Extension: ".db",
		}

		err := dbutils.OpenDatabase(dbFilePath)
		if err != nil {
			log.Fatalf("Fail open database: %s", err)
		}
		defer dbutils.CloseDatabase()

		err = dbstatements.InitDBStatements(dbutils.DB)
		if err != nil {
			log.Fatalf("Fail to prepare database statements: %s", err)
		}

		defer dbstatements.CloseDBStatements()
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
	db, err := sql.Open("sqlite3", "./pkg/db/socialNetwork.db")
	if err != nil {
		log.Fatalf("err %s", err)
	}
	// user := &dbmodels.User{
	// 	UserId:         "2",
	// 	IsLoggedIn:     1,
	// 	Email:          "user@test.com2",
	// 	HashedPassword: "hashed_password2",
	// 	FirstName:      "First2",
	// 	LastName:       "Last2",
	// 	DOB:            time.Now(),
	// 	AvatarPath:     "path/to/avatar2",
	// 	DisplayName:    "User2",
	// 	AboutMe:        "About me2",
	// }
	// addressOfValues := StructFieldValues(user)
	// for i, v := range addressOfValues {
	// 	fmt.Println(i, v)
	// }
	// log.Println(len(addressOfValues))
	// err = userdb.InsertUser(db, addressOfValues)
	// if err != nil {
	// 	log.Fatalf("err: %s", err)
	// }
	user, err := crud.SelectFromDatabase(db, "Users", "WHERE user_id = 2")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}
