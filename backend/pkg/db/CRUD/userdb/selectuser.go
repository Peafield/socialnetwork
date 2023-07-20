package userdb

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

/*
SelectUser returns a user from the database given a column value and a search value.

Firstly, validates that the column value is present in the database table, so that it can query the database properly.
Then, opens the database, prepares the statement, and finally queries the row before returning the user as a User struct.

Parameters:
- db (*sql.DB): An open database to access and interact with.
- columnValue (string): refers to the column name in the sql table.
- searchValue (string): what to search for in the relevant column.

Returns:
- dbmodels.User: a user struct containing all the details of the queried user.
- error: any error relating to selecting the user.

Errors:
- Returns an error if columnName does not exist within the table by checking the relevant struct.
- Returns an error if the database failed to open.
- Returns an error if preparing the statement failed.
- Returns an error if querying the rows fails.

Examples:
- Used when selecting a user, maybe to retrieve certain details or verify credentials.
*/
func SelectUser(db *sql.DB, table string, conditionStatement string) (interface{}, error) {
	var object dbmodels.User

	stm := "SELECT * FROM " + table + " " + conditionStatement

	// Use a prepared statement to prevent SQL injection.
	stmt, err := db.Prepare(stm)
	if err != nil {
		return object, fmt.Errorf("statement preparation failed when selecting: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Query()
	if err != nil {
		return object, fmt.Errorf("query execution failed when selecting: %w", err)
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(helpers.StructFieldAddress(&object)...)
		if err != nil {
			return object, fmt.Errorf("failed to scan user data: %w", err)
		}
	}
	return object, err
}
