package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

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
