package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
)

func UpdateFollowStatus(db *sql.DB, userId string, updateFollowerData map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["followee_id"] = updateFollowerData["followee_id"].(string)
	conditions["follower_id"] = userId

	immutableParameters := []string{"creation_date"}

	dataContainsImmutableParameter := helpers.MapKeyContains(updateFollowerData, immutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update follower immutable parameter")
	}

	err := crud.UpdateDatabaseRow(db, "Followers", conditions, updateFollowerData)
	if err != nil {
		return fmt.Errorf("failed to update follower data: %w", err)
	}
	return nil
}
