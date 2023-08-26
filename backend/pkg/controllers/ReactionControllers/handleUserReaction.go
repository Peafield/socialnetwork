package reactioncontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func updateReactionCounts(db *sql.DB, reactionOn, reactionType, postORCommentId string) error {
	var updateQuery *sql.Stmt
	switch {
	case reactionOn == "post" && reactionType == "like":
		updateQuery = dbstatements.UpdatePostIncreaseLikeStmt
	case reactionOn == "post" && reactionType == "dislike":
		updateQuery = dbstatements.UpdatePostIncreaseDislikeStmt
	case reactionOn == "comment" && reactionType == "like":
		updateQuery = dbstatements.UpdateCommentIncreaseLikeStmt
	case reactionOn == "comment" && reactionType == "dislike":
		updateQuery = dbstatements.UpdateCommentIncreaseDislikeStmt
	}
	return crud.InteractWithDatabase(db, updateQuery, []interface{}{postORCommentId})
}

func handleExistingReaction(db *sql.DB, existingReaction *dbmodels.Reaction, reactionType, userID, postORCommentId, reactionOn string) error {
	var deleteQuery, updateQuery, decreaseQuery, increaseQuery *sql.Stmt
	// Initialize delete and update queries based on reactionOn value
	if reactionOn == "post" {
		deleteQuery = dbstatements.DeletePostReaction
		updateQuery = dbstatements.UpdatePostReaction
		if existingReaction.Reaction == "like" {
			decreaseQuery = dbstatements.UpdatePostDecreaseLikesStmt
		} else {
			decreaseQuery = dbstatements.UpdatePostDecreaseDislikesStmt
		}
		if reactionType == "like" {
			increaseQuery = dbstatements.UpdatePostIncreaseLikeStmt
		} else {
			increaseQuery = dbstatements.UpdatePostIncreaseDislikeStmt
		}
	} else {
		deleteQuery = dbstatements.DeleteCommentReaction
		updateQuery = dbstatements.UpdateCommentReaction
		if existingReaction.Reaction == "like" {
			decreaseQuery = dbstatements.UpdateCommentDecreaseLikesStmt
		} else {
			decreaseQuery = dbstatements.UpdateCommentDecreaseDislikesStmt
		}
		if reactionType == "like" {
			increaseQuery = dbstatements.UpdateCommentIncreaseLikeStmt
		} else {
			increaseQuery = dbstatements.UpdateCommentIncreaseDislikeStmt
		}
	}

	if existingReaction.Reaction == reactionType {
		// Delete if user has un(dis)liked
		err := crud.InteractWithDatabase(db, deleteQuery, []interface{}{userID, postORCommentId})
		if err != nil {
			return fmt.Errorf("failed to delete reaction in database: %w", err)
		}

		// Decrease the like or dislike count
		err = crud.InteractWithDatabase(db, decreaseQuery, []interface{}{postORCommentId})
		if err != nil {
			return fmt.Errorf("failed to decrease likes/dislikes in database: %w", err)
		}
	} else {
		// Update if user changed from like to dislike or vice-versa
		err := crud.InteractWithDatabase(db, updateQuery, []interface{}{reactionType, userID, postORCommentId})
		if err != nil {
			return fmt.Errorf("failed to update reaction in database: %w", err)
		}

		// Increase the like or dislike count for the new reaction type
		err = crud.InteractWithDatabase(db, increaseQuery, []interface{}{postORCommentId})
		if err != nil {
			return fmt.Errorf("failed to increase likes/dislikes in database: %w", err)
		}

		// Decrease the like or dislike count for the old reaction type
		err = crud.InteractWithDatabase(db, decreaseQuery, []interface{}{postORCommentId})
		if err != nil {
			return fmt.Errorf("failed to decrease likes/dislikes in database: %w", err)
		}
	}

	return nil
}

func HandleUserReaction(db *sql.DB, userID string, newReactionData map[string]interface{}) error {
	selectArgs := make([]interface{}, 2)
	args := make([]interface{}, 4)

	selectArgs[0] = userID
	args[0] = userID

	postORCommentId, ok := newReactionData["reactionOnId"].(string)
	if !ok {
		return fmt.Errorf("failed to assert whether post or comment")
	}

	reactionOn, ok := newReactionData["reactionOn"].(string)
	if !ok {
		return fmt.Errorf("reactionOn is not a string")
	}

	reactionType, ok := newReactionData["reactionType"].(string)
	if !ok {
		return fmt.Errorf("reactionType is not a string")
	}

	args[3] = reactionType

	var selectQuery *sql.Stmt

	// Initialize statements based on reactionOn value
	switch reactionOn {
	case "post":
		selectArgs[1] = postORCommentId
		args[1] = postORCommentId
		selectQuery = dbstatements.SelectPostReactionStmt
	case "comment":
		selectArgs[1] = postORCommentId
		args[2] = postORCommentId
		selectQuery = dbstatements.SelectCommentReactionStmt
	default:
		return fmt.Errorf("invalid reactionOn value: %s", reactionOn)
	}

	// Query and handle existing reactions
	reactionExists, err := crud.SelectFromDatabase(db, "Reactions", selectQuery, selectArgs)
	if !errors.Is(err, errorhandling.ErrNoResultsFound) && err != nil {
		return fmt.Errorf("failed to select reaction from database: %w", err)
	}

	if len(reactionExists) > 0 {
		existingReaction, ok := reactionExists[0].(*dbmodels.Reaction)
		if ok {
			err := handleExistingReaction(db, existingReaction, reactionType, userID, postORCommentId, reactionOn)
			if err != nil {
				return err
			}
		}
	} else {
		err = crud.InteractWithDatabase(db, dbstatements.InsertReactionsStmt, args)
		if err != nil {
			return fmt.Errorf("failed to insert reaction into database: %w", err)
		}
		err = updateReactionCounts(db, reactionOn, reactionType, postORCommentId)
		if err != nil {
			return err
		}
	}

	return nil
}
