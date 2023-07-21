package main

import (
	"flag"
	"log"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
)

// YOU MUST CALLED --dbopen WHEN STARTING THE SERVER TO OPEN THE DATABASE

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
			log.Println("Missing database name")
		}

		dbFilePath := &helpermodels.FilePathComponents{
			Directory: DATABASE_FILE_PATH,
			FileName:  dbName,
			Extension: ".db",
		}

		err := dbutils.OpenDatabase(dbFilePath)
		if err != nil {
			log.Printf("Failed open database: %s", err)
		} else {
			defer dbutils.CloseDatabase()
		}

		err = dbstatements.InitDBStatements(dbutils.DB)
		if err != nil {
			log.Printf("Failed to prepare database statements: %s", err)
		} else {
			defer dbstatements.CloseDBStatements()
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

	// FOR TESTING PURPOSES MUST DELETE
	// user := &dbmodels.User{
	// 	UserId:         "4",
	// 	IsLoggedIn:     1,
	// 	Email:          "user@test.com4",
	// 	HashedPassword: "hashed_password4",
	// 	FirstName:      "First4",
	// 	LastName:       "Last4",
	// 	DOB:            time.Now(),
	// 	AvatarPath:     "path/to/avatar4",
	// 	DisplayName:    "User4",
	// 	AboutMe:        "About me4",
	// }
	// values := StructFieldValues(user)
	// err := crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertUserStmt, values)
	// if err != nil {
	// 	log.Fatalf("failed to insert user data into db: %s", err)
	// }
	// user := &dbmodels.Post{
	// 	PostId:           "1",
	// 	GroupId:          "1",
	// 	CreatorId:        "1",
	// 	Title:            "TEST1",
	// 	ImagePath:        "path/to/image",
	// 	Content:          "A whole bunch of nonsense",
	// 	PrivacyLevel:     0,
	// 	AllowedFollowers: "ted, jill, andrew",
	// 	Likes:            100,
	// 	Dislikes:         100000,
	// }
	// values := helpers.StructFieldValues(user)
	// err := crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertPostStmt, values)
	// if err != nil {
	// 	log.Fatalf("err %s", err)
	// }
	err := crud.DeleteFromDatabase(dbutils.DB, "Posts", "post_id", "1")
	if err != nil {
		log.Fatalf("delete failed: %s", err)
	}
}
