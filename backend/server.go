package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	db "socialnetwork/pkg/db/mocking"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
	"socialnetwork/pkg/routehandlers"
	"socialnetwork/pkg/websocket"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// YOU MUST CALL --dbopen WHEN STARTING THE SERVER TO OPEN THE DATABASE

const DATABASE_FILE_PATH = "./pkg/db/"
const MIGRATIONS_FILE_PATH = "./pkg/db/migrations"

func main() {
	/*FLAGS*/
	dbinit := flag.Bool("dbinit", false, "Initialises a database")
	dbopen := flag.Bool("dbopen", false, "Opens a database and prepares database statements")
	dbOpenFlag := os.Getenv("DBOPEN_FLAG")
	dbup := flag.Bool("dbup", false, "Migrate database changes up")
	dbdown := flag.Bool("dbdown", false, "Migrate database changes down")
	dbmock := flag.Bool("dbmock", false, "Creates mock data for the database. You must run dbopen before running dbmock")

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
		return
	}

	if *dbopen || dbOpenFlag != "" {
		var dbName string
		if dbOpenFlag != "" {
			dbName = dbOpenFlag
		} else {
			dbName = flag.Arg(0)
		}
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

		reset := func() {
			err := controllers.SignOutAllUsers(dbutils.DB)
			if err != nil {
				log.Fatalf("error signing out all users: %s", err)
			}
		}
		reset()

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
		return
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
		return
	}

	if *dbmock {
		err := db.CreateFakeUsers(dbutils.DB)
		if err != nil {
			log.Fatalf("someting went wrong creating fakes: %s", err)
		}
		err = db.CreateFakePosts(dbutils.DB)
		if err != nil {
			log.Fatalf("someting went wrong creating fakes: %s", err)
		}
		err = db.CreateFakeComments(dbutils.DB)
		if err != nil {
			log.Fatalf("something went wrong creating fakes: %s", err)
		}
		err = db.CreateFakeFollowers(dbutils.DB)
		if err != nil {
			log.Fatalf("something went wrong creating fakes: %s", err)
		}
		err = db.CreateFakeChats(dbutils.DB)
		if err != nil {
			log.Fatalf("something went wrong creating fakes: %s", err)
		} else {
			log.Println("Mocking Successful!")
		}
	}

	/*SERVER SETTINGS*/
	r := mux.NewRouter()
	hub := websocket.NewHub()
	go hub.Run()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// Attach CORS middleware to your router
	r.Use(handlers.CORS(originsOk, headersOk, methodsOk))

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/*AUTH ENDPOINTS*/
	r.Handle("/signup", middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.SignUpHandler)))
	r.Handle("/signin", middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.SignInHandler)))
	r.Handle("/signout", middleware.ValidateTokenMiddleware(http.HandlerFunc(routehandlers.SignOutHandler)))

	/*END POINTS*/
	r.Handle("/user", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.UserHandler))))
	r.Handle("/post", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.PostHandler))))
	r.Handle("/comment", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.CommentHandler))))
	r.Handle("/reaction", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.ReactionHandler))))
	r.Handle("/follow", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.FollowerHandler))))
	r.Handle("/notification", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.NotificationHandler))))
	r.Handle("/group", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.GroupsHandler))))
	r.Handle("/groupmembers", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.GroupMembersHandler))))
	r.Handle("/event", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.GroupEventsHandler))))
	r.Handle("/eventattendees", middleware.ValidateTokenMiddleware(middleware.ParseAndValidateData(http.HandlerFunc(routehandlers.GroupEventAttendeesHandler))))

	/*WEBSOCKET*/
	r.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	}))

	/*LISTEN AND SERVER*/
	log.Printf("Server running on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
