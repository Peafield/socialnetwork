package main

import (
	"flag"
	"log"
	"net/http"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
	"socialnetwork/pkg/routehandlers"
	"time"

	"github.com/gorilla/mux"
)

// YOU MUST CALL --dbopen WHEN STARTING THE SERVER TO OPEN THE DATABASE

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

	/*SERVER SETTINGS*/
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Handle("/signup", middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.SignUpHandler))).Methods("POST")
	r.Handle("/signin", middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.SignInHandler))).Methods("POST")

	/*LISTEN AND SERVER*/
	log.Fatal(srv.ListenAndServe())
	// pword, _ := helpers.HashPassword("abc123")
	// user := &dbmodels.User{
	// 	UserId:         "10",
	// 	IsLoggedIn:     0,
	// 	Email:          "test@test101.com",
	// 	HashedPassword: pword,
	// 	FirstName:      fmt.Sprintf("F%v", time.Now().Unix()),
	// 	LastName:       fmt.Sprintf("L%v", time.Now().Unix()),
	// 	DOB:            time.Now(),
	// 	AvatarPath:     "/path/to/avatar",
	// 	DisplayName:    "testDisplayName1",
	// 	AboutMe:        fmt.Sprintf("Test time: %v", time.Now().Unix()),
	// }
	// values, _ := helpers.StructFieldValues(user)
	// err := crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertUserStmt, values)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
