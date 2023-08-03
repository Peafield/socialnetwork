package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

func SignOutAllUsers(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized, please run -dbopen")
	}
	affectedColumns := map[string]interface{}{}
	affectedColumns["is_logged_in"] = 0
	err := crud.UpdateDatabaseRow(db, "Users", map[string]interface{}{}, affectedColumns)
	if err != nil {
		return fmt.Errorf("failed to reset all user logged in status to 0: %w", err)
	}
	return nil
}
