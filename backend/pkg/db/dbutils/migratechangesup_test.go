package db_test

import (
	"errors"
	"log"
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"testing"

	"github.com/golang-migrate/migrate/v4"
)

//what do we need to test
//test validation of path whether success or fail

//Simulate an error for the initialization of "migrate"

// simulate an error for the Up migration

type MockMigrationInit struct{}
type MockMigrationUpDown struct{}

func (e *MockMigrationInit) New(sourceURL string, databaseURL string) (*migrate.Migrate, error) {
	return nil, errors.New("stfu")
}
func (e *MockMigrationUpDown) Up(m *migrate.Migrate) error   { return errors.New("up") }
func (e *MockMigrationUpDown) Down(m *migrate.Migrate) error { return errors.New("Down") }

func TestMigrateChangesUp(t *testing.T) {
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

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the test
			dbFilePath := &dbmodels.DatabaseFilePathComponents{
				DBName:    tc.dbName,
				Directory: tc.dbDirectory,
			}
			dberr := db.InitialiseDatabase(dbFilePath)
			if dberr != nil && tc.isValidPath {
				t.Errorf("db.CreateDatabase error: %s", dberr)
			}
			err := db.MigrateChangesUp(dbFilePath, "../migrations", tc.migrationConstructor, tc.migrateUpDown)

			// Assertions
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
