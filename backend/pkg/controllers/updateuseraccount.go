package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
)

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
