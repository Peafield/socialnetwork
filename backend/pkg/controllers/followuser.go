package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

/*
FollowUser
*/
func FollowUser(db *sql.DB, userId string, postFollowerData map[string]interface{}) error {
	follower := &dbmodels.Follower{}

	//set follower id
	follower.FollowerId = userId

	//set followee id
	followeeId, ok := postFollowerData["followee_id"].(string)
	if !ok {
		return fmt.Errorf("followee id is not a string")
	}
	follower.FolloweeId = followeeId

	//set following status (0 because it will be pending)
	follower.FollowingStatus = 0

	//set request pending (1 as request has now been sent)
	follower.RequestPending = 1

	//get follower struct values
	values, err := helpers.StructFieldValues(follower)
	if err != nil {
		return fmt.Errorf("failed to get follower struct values: %s", err)
	}

	err = crud.InsertIntoDatabase(db, dbstatements.InsertFollowersStmt, values)
	if err != nil {
		return fmt.Errorf("failed to insert follower into database, err: %s", err)
	}

	return nil
}
