package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

func NewPostSelectedFollower(db *sql.DB, userId string, newPostSelectedFollowerData map[string]interface{}) error {
	postSelectedFollower := &dbmodels.PostSelectedFollower{}

	//get post id
	postId, ok := newPostSelectedFollowerData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	postSelectedFollower.PostId = postId

	//get allowed follower id
	allowedFollowerId, ok := newPostSelectedFollowerData["allowed_follower_id"].(string)
	if !ok {
		return fmt.Errorf("allowed follower id is not a string")
	}
	postSelectedFollower.AllowedFollowerId = allowedFollowerId

	//get post selected follower struct values
	values, err := helpers.StructFieldValues(postSelectedFollower)
	if err != nil {
		return fmt.Errorf("failed to get post selected follower struct values: %s", err)
	}

	err = crud.InsertIntoDatabase(db, dbstatements.InsertPostsSelectedFollowerStmt, values)
	if err != nil {
		return fmt.Errorf("failed to insert post selected follower into database, err: %s", err)
	}

	return nil
}
