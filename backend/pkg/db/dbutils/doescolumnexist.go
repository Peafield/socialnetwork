package dbutils

import (
	"database/sql"
	"fmt"
)

/*
DoesColumnExist makes sure a column is defined within a certain database table.

Queries the table using PRAGMA to access non-table-data information such as the column names.

Parameters:
- db (*sql.DB): an open database
- tableName (string): the name of the relevant database table.
- columnName (string): the name of the column to check for.

Returns:
- bool for validation
- error if the field cannot be found in the struct

Example:
- Is used when interacting with the database to make sure the column is in the database
*/
func DoesColumnExist(db *sql.DB, tableName, columnName string) (bool, error) {
	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var cid int
		var name string
		var typeVar string
		var notNull int
		var dfltValue *string
		var pk int

		if err := rows.Scan(&cid, &name, &typeVar, &notNull, &dfltValue, &pk); err != nil {
			return false, err
		}
		if name == columnName {
			return true, nil
		}
	}

	return false, nil
}
