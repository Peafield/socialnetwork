package dbutils_test

import (
	"errors"
	"log"
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/helpermodels"
	"testing"
)

// Simulate migration changes up in a temporary database.
func TestMigrateChangesDown(t *testing.T) {
	tempPath := t.TempDir()
	log.Println(tempPath)

	testCases := []struct {
		name                 string
		dbName               string
		dbDirectory          string
		migrationConstructor dbmodels.MigrationConstructor
		migrateUpDown        dbmodels.MigrationUpdates
		isValidPath          bool
		migrationInitiliased bool
		migrationSucceeded   bool
		expectError          bool
	}{
		{
			name:                 "Everything works",
			dbName:               "mydatabase",
			dbDirectory:          tempPath,
			migrationConstructor: &dbmodels.NativeMigrate{},
			migrateUpDown:        &dbmodels.NativeMigrateUpdates{},
			isValidPath:          true,
			migrationInitiliased: true,
			migrationSucceeded:   true,
			expectError:          false,
		},
		{
			name:                 "Invalid path",
			dbName:               "mydatabase",
			dbDirectory:          "/invalid/directory",
			migrationConstructor: &dbmodels.NativeMigrate{},
			migrateUpDown:        &dbmodels.NativeMigrateUpdates{},
			isValidPath:          false,
			migrationInitiliased: false,
			migrationSucceeded:   false,
			expectError:          true,
		},
		{
			name:                 "Migration initialisation failure",
			dbName:               "mydatabase",
			dbDirectory:          tempPath,
			migrationConstructor: &MockMigrationInit{},
			migrateUpDown:        &dbmodels.NativeMigrateUpdates{},
			isValidPath:          true,
			migrationInitiliased: false,
			migrationSucceeded:   false,
			expectError:          true,
		},
		{
			name:                 "Migration failure",
			dbName:               "mydatabase",
			dbDirectory:          tempPath,
			migrationConstructor: &dbmodels.NativeMigrate{},
			migrateUpDown:        &MockMigrationUpDown{},
			isValidPath:          true,
			migrationInitiliased: true,
			migrationSucceeded:   false,
			expectError:          true,
		},
	}

	var ErrDatabaseExists = errors.New("database already exists")
	var ErrInvalidPath = errors.New("invalid directory path")

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &helpermodels.FilePathComponents{
				FileName:  tc.dbName,
				Directory: tc.dbDirectory,
				Extension: ".db",
			}
			dberr := db.CreateDatabase(dbFilePath)
			if errors.Is(dberr, ErrInvalidPath) && tc.isValidPath {
				t.Errorf("db.InitialiseDatabase error: %s", dberr)
			}
			// Continue onto migrations if the error is ErrDatabaseExists
			if dberr == nil || errors.Is(dberr, ErrDatabaseExists) {
				err := db.MigrateChangesDown(dbFilePath, "../migrations", tc.migrationConstructor, tc.migrateUpDown)
				// Assertions
				if tc.expectError && err == nil {
					t.Error("Expected an error, but got nil")
				} else if !tc.expectError && err != nil {
					t.Errorf("Unexpected error: %s", err)
				}
			}
		})
	}
}
