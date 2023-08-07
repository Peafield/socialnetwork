package crud

import (
	"database/sql"
	"fmt"
)

/*
DeleteFromDatabase deletes a row from a table in a database by a specific value.

The function will taken an open database, a table to delete from, and a value which will
indicate which row should be deleted. It creates and prepares a statement to exectute and then
checks if any rows have been sucessfully deleted.

Parameters:
  - db (*sql.DB): an open connection to a sql database.
  - table (string): the database table to be deleted from.
  - value (string): the value by which a row should be deleted.

Returns:
  - error: An error will be returned if the statement fails to be prepared or executed, if it fails to retrieve row data,
    and if no rows have been affected during insertion.

Example:
  - A user could be deleted from a the Users table by inserting the user's id as the value.
*/
func DeleteFromDatabase(db *sql.DB, query string, args ...interface{}) error {
	result, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected during deletion")
	}

	return nil
}
