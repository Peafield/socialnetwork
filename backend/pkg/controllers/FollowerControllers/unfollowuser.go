package followercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

/*
UnfollowUser is used to remove the record of a specific follower.

It defines the search conditions using the passed in data then call the DeleteFromDatabase function.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - deleteFollowerData (map[string]interface{}): data about the follower to delete.

Errors:
  - failure to delete the follower record from the database.
*/
func UnfollowUser(db *sql.DB, userId string, deleteFollowerData map[string]interface{}) error {
	var args []interface{}

	if followeeId, ok := deleteFollowerData["followee_id"].(string); ok {
		args = append(args, followeeId)
	}
	if followerId, ok := deleteFollowerData["follower_id"].(string); ok {
		args = append(args, followerId)
	}

	query := fmt.Sprintf("DELETE FROM Followers WHERE followee_id = ? AND follower_id = ?")
	deleteFollowerStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare delete follower statement: %w", err)
	}
	defer deleteFollowerStatment.Close()

	//delete
	err = crud.InteractWithDatabase(db, deleteFollowerStatment, args)
	if err != nil {
		return fmt.Errorf("failed to delete follower data: %w", err)
	}
	return nil
}
