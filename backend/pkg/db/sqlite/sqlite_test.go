package db_test

import (
	"database/sql"
	"errors"
	"os"
	db "socialnetwork/pkg/db/sqlite"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
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
			dbFilePath := &dbmodels.BasicDatabaseInit{
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

type MockFileCreator struct{}

func (f *MockFileCreator) Create(name string) (*os.File, error) {
	return nil, errors.New("cannot create file")
}

type MockSQLDBOpener struct{}

func (o *MockSQLDBOpener) Open(driveName, dataSourceName string) (*sql.DB, error) {
	return nil, errors.New("cannot open database")
}

func TestCreateDatabase(t *testing.T) {
	tempPath := t.TempDir()

	testCases := []struct {
		name          string
		dbName        string
		dbDirectory   string
		fileCreator   helpermodels.FileCreator
		dbOpener      dbmodels.DBOpener
		dbFileCreated bool
		dbFileOpened  bool
		expectError   bool
	}{
		{
			name:          "Database created and opened",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			fileCreator:   &helpermodels.OSFileCreator{},
			dbOpener:      &dbmodels.SQLDBOpener{},
			dbFileCreated: true,
			dbFileOpened:  true,
			expectError:   false,
		},
		{
			name:          "Database not created",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			fileCreator:   &MockFileCreator{},
			dbOpener:      &dbmodels.SQLDBOpener{},
			dbFileCreated: false,
			dbFileOpened:  false,
			expectError:   true,
		},
		{
			name:          "Database created but not opened",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			fileCreator:   &helpermodels.OSFileCreator{},
			dbOpener:      &MockSQLDBOpener{},
			dbFileCreated: true,
			dbFileOpened:  false,
			expectError:   true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &dbmodels.BasicDatabaseInit{
				DBName:    tc.dbName,
				Directory: tc.dbDirectory,
			}
			err := db.CreateDatabase(dbFilePath.Directory, dbFilePath.DBName, tc.fileCreator, tc.dbOpener)

			// Assertions
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
