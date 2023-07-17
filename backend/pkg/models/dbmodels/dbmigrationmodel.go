package dbmodels

import (
	"github.com/golang-migrate/migrate/v4"
)

type MigrationConstructor interface {
	New(sourceURL string, databaseURL string) (*migrate.Migrate, error)
}

type MigrationUpdates interface {
	Up(m *migrate.Migrate) error
	Down(m *migrate.Migrate) error
}

type NativeMigrate struct{}

func (m *NativeMigrate) New(sourceURL string, databaseURL string) (*migrate.Migrate, error) {
	return migrate.New(sourceURL, databaseURL)
}

type NativeMigrateUpdates struct{}

func (mup *NativeMigrateUpdates) Up(m *migrate.Migrate) error {
	return m.Up()
}

func (mup *NativeMigrateUpdates) Down(m *migrate.Migrate) error {
	return m.Down()
}
