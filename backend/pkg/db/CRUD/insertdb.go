package crud

import (
	"database/sql"
	"fmt"
)

/*
InsertIntoDatabase inserts any given data into a specific database.

The function will take an open database, an open statement and an interface of values
and insert them into the database.

Parameters:
  - db (*sql.DB): an open connection to a sql database.
  - statement (*sql.Stmt): a prepared database statement for insertion.
  - values ([]interface{}: an interface of values to be inserted.

Returns:
  - error: An error will be returned if the statement fails to execute, if it fails to retrieve row data,
    and if no rows have been affected during insertion.
*/
func InsertIntoDatabase(db *sql.DB, statement *sql.Stmt, values []interface{}) error {
	result, err := statement.Exec(values...)
	if err != nil {
		return fmt.Errorf("failed to exectute statement: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows: %s", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("now rows affected during insertion: %s", err)
	}

	return nil
}
