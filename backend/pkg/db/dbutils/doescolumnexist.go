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

/*
TableColumnNames returns a slice of string representing all of the columns from a given DB table.
this function can be used, despite a table undergoing DB migration.

Parameters:
- db (*sql.DB): an open database
- tableName (string): the name of the relevant database table.

Returns:
- []string: every single column found in the table
- error: any errors encountered during the query process

Example:
- To prevent the usage of non existing columns in a table.
*/
func TableColumnNames(db *sql.DB, tableName string) ([]string, error) {
	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next() {
		var name string

		if err := rows.Scan(nil, &name, nil, nil, nil, nil); err != nil {
			return nil, err
		}
		result = append(result, name)
	}

	return result, nil
}
