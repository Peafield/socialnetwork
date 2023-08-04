package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

/*
DeletePostSelectedFollower is used to remove the record of a specific post's selected follower.

It defines the search conditions using the passed in data then call the DeleteFromDatabase function.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - deletePostSelectedFollowerData (map[string]interface{}): data about the post selected follower to delete.

Errors:
  - failure to delete the post selected follower record from the database.
*/
func DeletePostSelectedFollower(db *sql.DB, userId string, deletePostSelectedFollowerData map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["post_id"] = deletePostSelectedFollowerData["post_id"]
	conditions["allowed_follower_id"] = deletePostSelectedFollowerData["allowed_follower_id"]

	err := crud.DeleteFromDatabase(db, "Posts_Selected_Followers", conditions)
	if err != nil {
		return fmt.Errorf("failed to delete post selected follower from database: %w", err)
	}
	return nil
}
