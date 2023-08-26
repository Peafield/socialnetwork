package reactioncontrollers

// func SelectReactions(db *sql.DB, reactionData map[string]interface{}) (*dbmodels.Reactions, error) {
// 	var reactionType, reaction string

// 	reaction, ok := reactionData["post_id"].(string)
// 	if ok {
// 		reactionType = "post_id"
// 	} else {
// 		reaction, ok = reactionData["commment_id"].(string)
// 		if ok {
// 			reactionType = "comment_id"
// 		}
// 	}

// 	query := `
// 		SELECT * FROM Reactions
// 		WHERE ` + reactionType + ` = ?
// 		`
// 	queryValues := []interface{}{
// 		reaction,
// 	}
// 	reactionsData, err := crud.SelectFromDatabase(db, "Reactions", query, queryValues)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to select reactions from database: %w", err)
// 	}

// 	reactions := &dbmodels.Reactions{}
// 	for _, v := range reactionsData {
// 		if reaction, ok := v.(dbmodels.Reaction); ok {
// 			reactions.Reactions = append(reactions.Reactions, reaction)
// 		} else {
// 			return nil, fmt.Errorf("failed to assert reaction data")
// 		}
// 	}
// 	return reactions, nil
// }
