package db

import (
	"database/sql"
	"fmt"
)

// should store the immutable columns in a map

func UpdateUserInfo(db *sql.DB, userId string, columnName string, value interface{}) error {
	//maybe check the fields first before accessing them

	Query := fmt.Sprintf(`UPDATE Users SET %s = %v WHERE user_id = %s;`, columnName, value, userId)
	statement, err := db.Prepare(Query)

	if err != nil {
		return fmt.Errorf("failed to prepare update user statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec()
	fmt.Println(result)
	if err != nil {
		return fmt.Errorf("failed to execute update user statement: %w", err)
	}
	//stick the result to a user

	return nil
}
