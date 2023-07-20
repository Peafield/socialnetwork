package crud

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/helpers"
)

func SelectFromDatabase(db *sql.DB, table string, conditionStatement string) (interface{}, error) {
	object := helpers.DecideStructType(table)

	stm := "SELECT * FROM " + table + " " + conditionStatement

	// Use a prepared statement to prevent SQL injection.
	stmt, err := db.Prepare(stm)
	if err != nil {
		return object, fmt.Errorf("failed to prepare select statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Query()
	if err != nil {
		return object, fmt.Errorf("failed to execute select query: %w", err)
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(helpers.StructFieldAddress(object)...)
		if err != nil {
			return object, fmt.Errorf("failed to scan user data: %w", err)
		}
	}
	return object, err
}
