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
	var table string
	var affectedColumn string
	var affectedValue string
	var reactionColumn string

	postId, ok := updateReactionData["post_id"].(string)
	if ok {
		conditions["post_id"] = postId
		table = "Posts"
		affectedColumn = "post_id"
		affectedValue = postId

	}

	commentId, ok := updateReactionData["comment_id"].(string)
	if ok {
		conditions["comment_id"] = commentId
		table = "Comments"
		affectedColumn = "comment_id"
		affectedValue = commentId
	}

	reactionInt, ok := updateReactionData["reaction"].(int)
	if ok {
		if reactionInt > 0 {
			reactionColumn = "likes"
		} else {
			reactionColumn = "dislikes"
		}
	}

	immutableParameters := []string{"user_id", "post_id", "comment_id", "creation_date"}

	dataContainsImmutableParameter := helpers.MapKeyContains(updateReactionData, immutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update reaction immutable parameter")
	}

	err := crud.UpdateDatabaseRow(db, "Reactions", conditions, updateReactionData)
	if err != nil {
		return fmt.Errorf("failed to update reaction data: %w", err)
	}

	postOrCommentConditions := map[string]interface{}{}
	postOrCommentConditions[affectedColumn] = affectedValue

	updatePostOrCommentData := map[string]interface{}{}
	updatePostOrCommentData[reactionColumn] = reactionColumn + "+ 1"

	err = crud.UpdateDatabaseRow(db, table, postOrCommentConditions, updatePostOrCommentData)
	if err != nil {
		return fmt.Errorf("failed to update reaction data: %w", err)
	}
	return nil
}
