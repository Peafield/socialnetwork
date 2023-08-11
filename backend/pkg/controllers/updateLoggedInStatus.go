package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"strings"
)

func UpdateLoggedInStatus(db *sql.DB, userId string, loggedInStatus int) error {
	var columns []string
	var args []interface{}

	if loggedInStatus > 1 || loggedInStatus < 0 {
		return fmt.Errorf("logged in status is not a valid value")
	}

	columns = append(columns, "is_logged_in = ?")
	args = append(args, loggedInStatus)
	args = append(args, userId)

	query := fmt.Sprintf("UPDATE Users SET %s WHERE user_id = ?", strings.Join(columns, ", "))
	updateLoggedInStatusStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update logged in status: %w", err)
	}
	defer updateLoggedInStatusStatment.Close()

	err = crud.InteractWithDatabase(db, updateLoggedInStatusStatment, args)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update logged in status: %w", err)
	}
	return nil
}
