package dbmodels

import "database/sql"

// DBOpener is an interface that provides methods to open a database.
type DBOpener interface {
	Open(driveName, dataSourceName string) (*sql.DB, error)
}

// SQLDBOpener is a struct that holds a drive name and a data source.
type SQLDBOpener struct {
	DriveName      string
	DataSourceName string
}

// Open returns a sqlite 3 database and an error.
func (o *SQLDBOpener) Open(driveName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driveName, dataSourceName)
}
