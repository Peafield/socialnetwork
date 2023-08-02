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
func UpdateDatabaseRow(DB *sql.DB, TableName string, Conditions map[string]interface{}, AffectedColumns map[string]interface{}) error {
	//maybe check the fields first before accessing them

	//seperate the keys and the values respectively as a string and an []interface{}
	setStatement, setValues := dbutils.UpdateSetConstructor(AffectedColumns)
	updateStatement, updatedValues := dbutils.ConditionStatementConstructor(Conditions)

	//assemble the final form of the statement
	finalStatement := fmt.Sprintf(`UPDATE %s %s %s;`, TableName, setStatement, updateStatement)

	statement, err := DB.Prepare(finalStatement)

	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %w", err)
	}

	defer statement.Close()

	//combine the values from both maps to pass in the "Exec()" method
	combinedValues := append(setValues, updatedValues...)
	result, err := statement.Exec(combinedValues...)

	if err != nil {
		return fmt.Errorf("failed to execute update statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows for update: %s", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected in update: %s", err)
	}

	return nil
}
