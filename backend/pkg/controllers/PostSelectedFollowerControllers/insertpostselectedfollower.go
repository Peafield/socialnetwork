package postselectedfollowercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

/*
NewPostSelectedFollower creates a PostSelectedFollower struct, fills out the properties using the passed in data,
and then gets the values as a slice of interface{} to then insert into the database.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current user id.
  - postPostSelectedFollowerData (map[string]interface{}): post selected follower data from the request.

Errors:
  - if the asserted values are not the right type.
  - failure to get the post selected follower values from the struct.
  - failure to insert the post selected follower record into the database.
*/
func NewPostSelectedFollower(db *sql.DB, userId string, newPostSelectedFollowerData map[string]interface{}) error {
	args := []interface{}{}

	//get post id
	postId, ok := newPostSelectedFollowerData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	args = append(args, postId)

	//get allowed follower id
	allowedFollowerId, ok := newPostSelectedFollowerData["allowed_follower_id"].(string)
	if !ok {
		return fmt.Errorf("allowed follower id is not a string")
	}
	args = append(args, allowedFollowerId)

	err := crud.InteractWithDatabase(db, dbstatements.InsertPostsSelectedFollowerStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post selected follower into database, err: %s", err)
	}

	return nil
}
