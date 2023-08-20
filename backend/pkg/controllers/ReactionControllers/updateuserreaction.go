package reactioncontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"strings"
)

func UpdateUserPostOrCommentReaction(db *sql.DB, userID string, updateReactionData map[string]interface{}) error {
	var columns []string
	var args []interface{}
	var updateTable string
	var updateValues string
	var updateColumn string

	if likes, ok := updateReactionData["likes"].(int); ok {
		columns = append(columns, "likes = ?")
		args = append(args, likes)
		updateValues = "likes = likes + 1"
	}
	if dislikes, ok := updateReactionData["dislikes"].(int); ok {
		columns = append(columns, "dislikes = ?")
		args = append(args, dislikes)
		updateValues = "dislikes = dislikes + 1"
	}
	if postId, ok := updateReactionData["post_id"].(string); ok {
		updateTable = "Posts"
		updateColumn = "post_id = ?"
		args = append(args, postId)

	}
	if commentId, ok := updateReactionData["comment_id"].(string); ok {
		updateTable = "Comments"
		updateColumn = "comment_id = ?"
		args = append(args, commentId)
	}

	query := fmt.Sprintf("UPDATE Reactions SET %s WHERE %s", strings.Join(columns, ", "), updateColumn)
	updateReactionsStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update reactions statement: %w", err)
	}
	defer updateReactionsStatment.Close()

	err = crud.InteractWithDatabase(db, updateReactionsStatment, args)
	if err != nil {
		return fmt.Errorf("failed to update reaction data: %w", err)
	}

	query = fmt.Sprintf("UPDATE %s SET %s WHERE %s", updateTable, updateValues, updateColumn)
	updatePostOrCommentStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update reactions statement: %w", err)
	}
	defer updatePostOrCommentStatment.Close()

	err = crud.InteractWithDatabase(db, updatePostOrCommentStatment, args)
	if err != nil {
		return fmt.Errorf("failed to update %s reaction data: %w", updateTable, err)
	}

	return nil
}
