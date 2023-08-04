package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

func UnfollowUser(db *sql.DB, userId string, deleteFollowerData map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["followee_id"] = deleteFollowerData["followee_id"].(string)
	conditions["follower_id"] = userId

	//delete
	err := crud.DeleteFromDatabase(db, "Followers", conditions)
	if err != nil {
		return fmt.Errorf("failed to delete follower data: %w", err)
	}
	return nil
}
