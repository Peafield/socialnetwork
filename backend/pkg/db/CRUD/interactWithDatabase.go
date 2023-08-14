package crud

import (
	"database/sql"
	"errors"
	"fmt"
	errorhandling "socialnetwork/pkg/errorHandling"
)

var ErrNoRowsAffected = errors.New("no rows affected")

// InteractWithDatabase executes a provided SQL statement on a given database connection
// using the provided arguments. It ensures that the statement affects at least one row
// and returns an error if any part of the execution process fails or if no rows are affected.
//
// Parameters:
//
//	db       - A pointer to an active database connection.
//	statment - A pointer to a prepared SQL statement ready for execution.
//	args     - A variadic set of arguments to pass to the statement for execution.
//
// Returns:
//   - nil if the statement is executed successfully and affects at least one row.
//   - An error if the execution fails, if retrieving the count of affected rows fails,
//     or if no rows are affected.
//
// Example:
//
//	stmt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	err = InteractWithDatabase(db, stmt, "NewName", 1)
//	if err != nil {
//	    log.Fatal(err)
//	}
func InteractWithDatabase(db *sql.DB, statement *sql.Stmt, args []interface{}) error {
	result, err := statement.Exec(args...)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return errorhandling.ErrNoRowsAffected
	}

	return nil
}
