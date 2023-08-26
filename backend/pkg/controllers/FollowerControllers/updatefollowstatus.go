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
func UpdateFollowStatus(db *sql.DB, userId string, updateFollowerData map[string]interface{}) error {
	var args []interface{}
	var query *sql.Stmt

	if followingStatus, ok := updateFollowerData["following_status"].(string); ok {
		query = dbstatements.UpdateFollowingStatusStmt
		args = append(args, followingStatus)
	}
	if requestPendingStatus, ok := updateFollowerData["request_pending"].(string); ok {
		query = dbstatements.UpdateRequestPendingStmt
		args = append(args, requestPendingStatus)
	}
	if followeeId, ok := updateFollowerData["followee_id"].(string); ok {
		args = append(args, followeeId)
	}
	if followerId, ok := updateFollowerData["follower_id"].(string); ok {
		args = append(args, followerId)
	}

	err := crud.InteractWithDatabase(db, query, args)
	if err != nil {
		return fmt.Errorf("failed to update post data: %w", err)
	}
	return nil
}
