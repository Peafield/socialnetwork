package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

func InsertUserReaction(db *sql.DB, userID string, newReactionData map[string]interface{}) error {
	var reaction *dbmodels.Reaction

	reaction.UserId = userID

	postID, ok := newReactionData["post_id"].(string)
	if ok {
		reaction.PostId = postID
	}

	commentID, ok := newReactionData["commment_id"].(string)
	if ok {
		reaction.CommentId = commentID
	}

	reaction.Reaction = newReactionData["reaction"].(int)

	values, err := helpers.StructFieldValues(reaction)
	if err != nil {
		return fmt.Errorf("failed to get reaction struct values: %w", err)
	}

	err = crud.InsertIntoDatabase(db, dbstatements.InsertReactionsStmt, values)
	if err != nil {
		return fmt.Errorf("failed to insert reaction into database: %w", err)
	}

	return nil
}
