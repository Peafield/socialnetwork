package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

/*
SignOutAllUsers makes sure that all users are signed out in the database.

The function accetps an open database. It then will map the is_logged_in column to be affected
and set to 0. This is the passed along with the database to the UpdateDatabaseRow function which will
do the updating in the database.

Parameters:
  - db(*sql.DB): an open database

Errors:
  - if the database connection is not initialised.
  - if there is an error updating the logged in status of users.

Example:
  - the function will be used on server restart.
*/
func SignOutAllUsers(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialised, please run -dbopen")
	}
	conditions := map[string]interface{}{}
	conditions["is_logged_in"] = 1
	affectedColumns := map[string]interface{}{}
	affectedColumns["is_logged_in"] = 0
	//conditions could be set to if logged in status is equal to 1 (for performance)
	err := crud.UpdateDatabaseRow(db, "Users", conditions, affectedColumns)
	if err != nil {
		return fmt.Errorf("failed to reset all user logged in status to 0: %w", err)
	}
	return nil
}
