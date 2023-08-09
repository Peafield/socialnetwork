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
	postId, ok := deletePostSelectedFollowerData["post_id"].(string)
	if !ok {
		return fmt.Errorf("postId is not a string")
	}
	allowedFollowerId, ok := deletePostSelectedFollowerData["allowed_follower_id"].(string)
	if !ok {
		return fmt.Errorf("allowedFollowerId is not a string")
	}

	args := []interface{}{}
	args = append(args, postId)
	args = append(args, allowedFollowerId)

	query := fmt.Sprintf("DELETE FROM Posts_Selected_Followers WHERE post_id = ? AND allowed_follower_id = ?")
	deletePostSelectedFollowerStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare delete Post selected follower statement: %w", err)
	}
	defer deletePostSelectedFollowerStatment.Close()

	//delete
	err = crud.InteractWithDatabase(db, deletePostSelectedFollowerStatment, args)
	if err != nil {
		return fmt.Errorf("failed to delete post selected follower from database: %w", err)
	}
	return nil
}
