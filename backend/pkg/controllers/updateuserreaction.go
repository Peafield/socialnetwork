package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
)

func UpdateUserReaction(db *sql.DB, userID string, updateReactionData map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["user_id"] = userID

	postId, ok := updateReactionData["post_id"]
	if ok {
		conditions["post_id"] = postId
	}

	commentId, ok := updateReactionData["comment_id"]
	if ok {
		conditions["comment_id"] = commentId
	}

	immutableParameters := []string{"user_id", "post_id", "comment_id", "creation_date"}

	dataContainsImmutableParameter := helpers.MapKeyContains(updateReactionData, immutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update reaction immutable parameter")
	}

	err := crud.UpdateDatabaseRow(db, "Posts", conditions, updateReactionData)
	if err != nil {
		return fmt.Errorf("failed to update reaction data: %w", err)
	}
	return nil
}
