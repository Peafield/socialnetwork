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

func SelectFollowerInfo(db *sql.DB, userId string, followeeId string) (*dbmodels.Follower, error) {
	followerData, err := crud.SelectFromDatabase(db, "Followers", dbstatements.SelectFollowerInfoStmt, []interface{}{userId, followeeId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select follower from database: %w", err)
	}

	if len(followerData) == 0 {
		return &dbmodels.Follower{}, nil
	}

	follower, ok := followerData[0].(*dbmodels.Follower)
	if !ok {
		return nil, fmt.Errorf("cannot assert follower type")
	}

	return follower, nil

}
