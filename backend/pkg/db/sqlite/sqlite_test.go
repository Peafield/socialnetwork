package db_test

import (
	db "socialnetwork/pkg/db/sqlite"
	"socialnetwork/pkg/models"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitialiseDatabase(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Define test cases
	testCases := []struct {
		name            string
		dbName          string
		dbDirectory     string
		isDBNameValid   bool
		isFilePathValid bool
		expectError     bool
	}{
		{
			name:            "Valid database name and directory",
			dbName:          "mydatabase",
			dbDirectory:     tmpDir,
			isDBNameValid:   true,
			isFilePathValid: true,
			expectError:     false,
		},
		{
			name:            "Invalid database name",
			dbName:          "my-database",
			dbDirectory:     tmpDir,
			isDBNameValid:   false,
			isFilePathValid: true,
			expectError:     true,
		},
		{
			name:            "Invalid database directory",
			dbName:          "mydatabase",
			dbDirectory:     "/invalid/directory",
			isDBNameValid:   true,
			isFilePathValid: false,
			expectError:     true,
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &models.BasicDatabaseInit{
				DBName:    tc.dbName,
				Directory: tc.dbDirectory,
			}
			err := db.InitialiseDatabase(dbFilePath)

			// Assertions
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}

func TestCreateDatabase(t *testing.T) {
	tempPath := t.TempDir()

	testCases := []struct {
		name          string
		dbName        string
		dbDirectory   string
		dbFileCreated bool
		dbFileOpened  bool
		expectError   bool
	}{
		{
			name:          "Database created and opened",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			dbFileCreated: true,
			dbFileOpened:  true,
			expectError:   false,
		},
		{
			name:          "Database not created",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			dbFileCreated: false,
			dbFileOpened:  false,
			expectError:   true,
		},
		{
			name:          "Database not created",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			dbFileCreated: true,
			dbFileOpened:  false,
			expectError:   true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &models.BasicDatabaseInit{
				DBName:    tc.dbName,
				Directory: tc.dbDirectory,
			}
			err := db.InitialiseDatabase(dbFilePath)

			// Assertions
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}

	// db.CreateDatabase(tempPath, testDBName)
	// filepath := path.Join(tempPath, testDBName)
	// _, err := os.Stat(filepath)
	// if os.IsNotExist(err) {
	// 	file, err := os.Create(filepath)
	// 	if err != nil {
	// 		t.Fatalf("file does not exist: %s", err)
	// 	}
	// 	file.Close()
	// }
	// db, err := sql.Open("sqlite3", filepath+"/.db?_foreign_keys=on")
	// if err != nil {
	// 	t.Fatalf("failed to open database: %s", err)
	// }
	// db.Close()
}
