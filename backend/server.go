package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"reflect"
	"socialnetwork/pkg/db/CRUD/userdb"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
	"time"
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
	db, err := sql.Open("sqlite3", "./pkg/db/socialNetwork.db")
	if err != nil {
		log.Fatalf("err %s", err)
	}
	user := &dbmodels.User{
		UserId:         "2",
		IsLoggedIn:     1,
		Email:          "user@test.com2",
		HashedPassword: "hashed_password2",
		FirstName:      "First2",
		LastName:       "Last2",
		DOB:            time.Now(),
		AvatarPath:     "path/to/avatar2",
		DisplayName:    "User2",
		AboutMe:        "About me2",
	}
	addressOfValues := StructFieldAddresses(user)
	for i, v := range addressOfValues {
		fmt.Println(i, v)
	}
	log.Println(len(addressOfValues))
	err = userdb.InsertUser(db, addressOfValues)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}

func StructFieldAddresses(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		if fieldType.Name != "CreationDate" {
			values = append(values, v.Field(i).Interface())
		}

	}
	return values
}
