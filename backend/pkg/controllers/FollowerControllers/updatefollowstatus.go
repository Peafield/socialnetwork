package followercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

/*
UpdateFollowStatus updates the follow status AND the request pending status of two users.

It creates the conditions map to find the record to refer to using the data passed through.  Then
creates a map of the parameters to be changed.  This can only be the following status and the request pending status,
these will be integers.  Once the information has been retrieved, the update can take place.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - updateFollowerData (map[string]interface{}): data about the follow status to update.

Errors:
  - failure to update the database

Example:
  - If a user want's to follow someone, they can request to follow.  Once this happens,
    you would use UpdateFollowStatus to change the request pending status from 0 to 1.
    If the receiving user accepts, you would change the request pending status from 1 back to 0,
    and the following status from 0 to 1.
*/
func UpdateFollowingStatus(db *sql.DB, userId string, updateFollowData map[string]interface{}) error {
	args := make([]interface{}, 3)

	followerId, ok := updateFollowData["follower_id"].(string)
	if !ok {
		return fmt.Errorf("follower_id is not a string or doesn't exist")
	}

	followingStatus, ok := updateFollowData["following_status"].(float64)
	if !ok {
		return fmt.Errorf("following_status is not a float64 or doesn't exist")
	}

	args[0] = int(followingStatus)
	args[1] = followerId
	args[2] = userId

	err := crud.InteractWithDatabase(db, dbstatements.UpdateFollowingStatusStmt, args)
	if err != nil {
		return fmt.Errorf("failed to update follower data: %w", err)
	}
	return nil
}
