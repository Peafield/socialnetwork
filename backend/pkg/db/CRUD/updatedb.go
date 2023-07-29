package crud

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/db/dbutils"
)

/*
UpdateTableRowInfo updates a particular row from a specified table by providing the columns to be changed and the conditions to produce the query.
the function prepares the query and returns a success or a failure if the row has been affected or not.

Parameters:
  - db (*sql.DB): an open connection to a sql database.
  - tableName (string): the database table to be deleted from.
  - Conditions (map[string]interface{}): maps the conditions to representing the desired row
  - AffectedColumns (map[string]interface{}): maps the columns to be updated.

Example:
  - A user would want to change his personal info (e.g nickname, profile picture ...).
*/
func UpdateDatabaseRow(db *sql.DB, tableName string, Conditions map[string]interface{}, AffectedColumns map[string]interface{}) error {
	//maybe check the fields first before accessing them

	var UpdatedValues string = dbutils.UpdateSetConstructor(AffectedColumns)
	var UpdatedConditions string = dbutils.ConditionStatementConstructor(Conditions)

	Query := fmt.Sprintf(`UPDATE %s %s %s;`, tableName, UpdatedValues, UpdatedConditions)

	statement, err := db.Prepare(Query)

	if err != nil {
		return fmt.Errorf("failed to prepare UpdateUserInfo's statement: %w", err)
	}

	defer statement.Close()

	result, err := statement.Exec()

	if err != nil {
		return fmt.Errorf("failed to execute UpdateUserInfo's statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows in UpdateUserInfo: %s", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected in UpdateUserInfo: %s", err)
	}

	return nil
}
