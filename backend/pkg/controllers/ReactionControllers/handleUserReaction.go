package reactioncontrollers

// func HandleUserReaction(db *sql.DB, userID string, newReactionData map[string]interface{}) error {
// 	selectArgs := make([]interface{}, 2)
// 	args := make([]interface{}, 4)
// 	var tableName, query string
// 	var deleteQuery, updateQuery *sql.Stmt

// 	selectArgs[0] = userID
// 	args[0] = userID

// 	postORCommentId, ok := newReactionData["reactionOnId"].(string)
// 	if !ok {
// 		return fmt.Errorf("failed to assert whether post or comment")
// 	}

// 	reactionOn, ok := newReactionData["reactionOn"].(string)
// 	if ok {
// 		selectArgs[1] = postORCommentId
// 		if reactionOn == "post" {
// 			args[1] = postORCommentId
// 			tableName = "Post"
// 			query = `
// 				SELECT * FROM REACTIONS
// 				WHERE user_id = ? AND post_id = ?
// 				`
// 			deleteQuery = dbstatements.DeletePostReaction
// 			updateQuery = dbstatements.UpdatePostReaction
// 		} else {
// 			args[2] = postORCommentId
// 			tableName = "Comment"
// 			query = `
// 				SELECT * FROM REACTIONS
// 				WHERE user_id = ? AND comment_id = ?
// 				`
// 			deleteQuery = dbstatements.DeleteCommentReaction
// 			updateQuery = dbstatements.UpdateCommentReaction
// 		}
// 	} else {
// 		return fmt.Errorf("reactionOn is not a string")
// 	}

// 	reactionType, ok := newReactionData["type"].(string)
// 	if ok {
// 		args[3] = reactionType
// 	} else {
// 		return fmt.Errorf("reactionType is not a string")
// 	}

// 	// check if the user interaction is already in the db. Delete if user has un(dis)liked, update if user
// 	// changed from like to dislike, insert if doesn't exist.
// 	reactionExists, err := crud.SelectFromDatabase(db, tableName, query, selectArgs)
// 	if err != nil {
// 		return fmt.Errorf("failed to select reaction from database: %w", err)
// 	}
// 	if len(reactionExists) > 0 {

// 		existingReaction, ok := reactionExists[0].(dbmodels.Reaction)
// 		if ok {
// 			if existingReaction.Reaction == "like" && reactionType == "like" || existingReaction.Reaction == "dislike" && reactionType == "dislike" {
// 				err = crud.InteractWithDatabase(db, deleteQuery, []interface{}{})
// 				if err != nil {
// 					return fmt.Errorf("failed to delete reaction in database: %w", err)
// 				}
// 			} else {
// 				err = crud.InteractWithDatabase(db, updateQuery, []interface{}{reactionType, userID, postORCommentId})
// 				return fmt.Errorf("failed to update reaction in database: %w", err)
// 			}
// 		}

// 	} else {
// 		err = crud.InteractWithDatabase(db, dbstatements.InsertReactionsStmt, args)
// 		if err != nil {
// 			return fmt.Errorf("failed to insert reaction into database: %w", err)
// 		}
// 	}

// 	return nil
// }
