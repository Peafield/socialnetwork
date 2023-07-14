package dbmodels

// DatabaseInit is an interface that provides methods to retrieve the database directory and name.
type DatabaseManager interface {
	GetDirectory() string
	GetDBName() string
}

// BasicDatabaseInit is a struct that holds the directory and name for a database.
type DatabaseFilePathComponents struct {
	Directory string
	DBName    string
}

// GetDirectory returns the directory of the database.
func (db *DatabaseFilePathComponents) GetDirectory() string {
	return db.Directory
}

// GetDBName returns the name of the database.
func (db *DatabaseFilePathComponents) GetDBName() string {
	return db.DBName
}
