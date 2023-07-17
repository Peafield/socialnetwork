package dbmodels

import "database/sql"

// DBOpener is an interface that provides methods to open a database.
type DBOpener interface {
	GetDriveName() string
	GetDataSourceName() string
	Open(driveName, dataSourceName string) (*sql.DB, error)
}

// SQLDBOpener is a struct that holds a drive name and a data source.
type SQLDBOpener struct {
	DriveName      string
	DataSourceName string
}

// GetDriveName returns the drive name, used in all functions that open the database
func (o *SQLDBOpener) GetDriveName() string {
	return o.DriveName
}

// GetDataSourceName returns the data source name, used in all functions that open the database
func (o *SQLDBOpener) GetDataSourceName() string {
	return o.DataSourceName
}

// Open returns a sqlite 3 database and an error.
func (o *SQLDBOpener) Open(driveName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driveName, dataSourceName)
}
