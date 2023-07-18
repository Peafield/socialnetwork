package dbutils_test

import (
	"database/sql"
	"errors"
	"os"
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/helpermodels"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

/*
the following function simulates the failed process of CREATING a file for testing purposes.
*/
type MockFileCreator struct{}

func (f *MockFileCreator) Create(name string) (*os.File, error) {
	return nil, errors.New("cannot create file")
}

/*
the following function simulates the failed process of OPENING a file for testing purposes.
*/
type MockSQLDBOpener struct {
	DriveName      string
	DataSourceName string
}

func (o *MockSQLDBOpener) GetDriveName() string {
	return o.DriveName
}

func (o *MockSQLDBOpener) GetDataSourceName() string {
	return o.DataSourceName
}

func (o *MockSQLDBOpener) Open(driveName, dataSourceName string) (*sql.DB, error) {
	return nil, errors.New("cannot open database")
}

func TestCreateDatabase(t *testing.T) {
	tempPath := t.TempDir()

	testCases := []struct {
		name          string
		dbName        string
		dbDirectory   string
		isValidPath   bool
		isFileCreated bool
		expectError   bool
	}{
		{
			name:          "Database valid and created",
			dbName:        "mydatabase",
			dbDirectory:   tempPath,
			isValidPath:   true,
			isFileCreated: true,
			expectError:   false,
		},
		{
			name:          "Invalid database name",
			dbName:        "@+][",
			dbDirectory:   tempPath,
			isValidPath:   false,
			isFileCreated: false,
			expectError:   true,
		},
		{
			name:          "invalid database directory",
			dbName:        "mydatabase1",
			dbDirectory:   "/invalid/directory",
			isValidPath:   false,
			isFileCreated: false,
			expectError:   true,
		},
		{
			name:          "Database already exists",
			dbName:        "mydatabase2",
			dbDirectory:   tempPath,
			isValidPath:   false,
			isFileCreated: false,
			expectError:   true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &helpermodels.FilePathComponents{
				Directory: tc.dbDirectory,
				FileName:  tc.dbName,
				Extension: ".db",
			}
			err := db.CreateDatabase(dbFilePath)
			if tc.name == "Database already exists" {
				err = db.CreateDatabase(dbFilePath)
			}
			// Assertions
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
