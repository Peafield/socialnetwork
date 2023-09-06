package followercontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectFollowersOfSpecificUser(db *sql.DB, follower_id string) (*dbmodels.Followers, error) {
	followersData, err := crud.SelectFromDatabase(db, "Followers", dbstatements.SelectFollowersOfUserStmt, []interface{}{follower_id})
	if err != nil && errors.Is(err, errorhandling.ErrNoRowsAffected) {
		return nil, fmt.Errorf("failed to select followers from database: %w", err)
	}

	followers := &dbmodels.Followers{}
	for _, v := range followersData {
		if follower, ok := v.(*dbmodels.Follower); ok {
			followers.Followers = append(followers.Followers, *follower)
		} else {
			return nil, fmt.Errorf("failed to assert follower data")
		}
	}

	return followers, nil
}
