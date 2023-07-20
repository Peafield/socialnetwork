package crud

import (
	"database/sql"
	"fmt"
)

/**/
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
