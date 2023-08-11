package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
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
	err := crud.InteractWithDatabase(db, dbstatements.UpdateAllUsersToSignedOut, []interface{}{})
	if err != nil {
		if errors.Is(err, errorhandling.ErrNoRowsAffected) {
			return nil
		} else {
			return fmt.Errorf("failed to reset all user logged in status to 0: %w", err)
		}
	}
	return nil
}
