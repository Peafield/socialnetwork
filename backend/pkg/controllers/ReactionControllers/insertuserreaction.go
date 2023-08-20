package reactioncontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func InsertUserReaction(db *sql.DB, userID string, newReactionData map[string]interface{}) error {
	args := make([]interface{}, 4)

	args[0] = userID

	rowIdType, ok := newReactionData["postORCommentId"].(string)
	if !ok {
		return fmt.Errorf("failed to assert whether post or comment")
	}

	reactionOn, ok := newReactionData["reactionOn"].(string)
	if ok {
		if reactionOn == "post" {
			args[1] = rowIdType
		} else {
			args[2] = rowIdType
		}
	} else {
		return fmt.Errorf("reactionOn is not a string")
	}

	reactionType, ok := newReactionData["type"].(string)
	if ok {
		args[3] = reactionType
	} else {
		return fmt.Errorf("reactionType is not a string")
	}

	err := crud.InteractWithDatabase(db, dbstatements.InsertReactionsStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert reaction into database: %w", err)
	}

	return nil
}
