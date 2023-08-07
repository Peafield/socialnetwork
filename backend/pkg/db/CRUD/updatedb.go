package crud

import (
	"database/sql"
	"fmt"
)

/*
UpdateTableRowInfo updates a particular row from a specified table by providing the columns to be changed and the conditions to produce the query.
the function prepares the query and returns a success or a failure if the row has been affected or not.

Parameters:
  - db (*sql.DB): an open connection to a sql database.
  - tableName (string): the database table to be deleted from.
  - conditions (map[string]interface{}): maps the conditions to representing the desired row
  - affectedColumns (map[string]interface{}): maps the columns to be updated.

Example:
  - A user would want to change his personal info (e.g nickname, profile picture ...).
*/
func UpdateDatabaseRow(db *sql.DB, query string, args ...interface{}) error {
	result, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows for update: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected in update")
	}

	return nil
}
