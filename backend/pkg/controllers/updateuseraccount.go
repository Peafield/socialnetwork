package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
)

/*
UpdateUserAccount updates the current users details.

It creates the conditions map to find the record to refer to using the data passed through.  Then
checks to make sure no immutable values are being passed through.  Then the update can take place.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - updateUserData (map[string]interface{}): data about the user account to update.

Errors:
  - an immutable property was found in the updateUserData map.
  - failure to update the database
*/
func UpdateUserAccount(db *sql.DB, userId string, updateUserData map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["user_id"] = userId

	//make sure immutable parameters are not trying to be changed
	immutableParameters := []string{"user_id", "is_logged_in", "creation_date"}

	dataContainsImmutableParameter := helpers.MapKeyContains(updateUserData, immutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update user immutable parameter")
	}

	//update user
	err := crud.UpdateDatabaseRow(db, "Users", conditions, updateUserData)
	if err != nil {
		return fmt.Errorf("error updating user, err: %s", err)
	}

	return nil
}
