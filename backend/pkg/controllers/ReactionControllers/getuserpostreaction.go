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

func GetUserPostReaction(db *sql.DB, userId string, postId string) (*dbmodels.Reaction, error) {
	reactionData, err := crud.SelectFromDatabase(db, "Reactions", dbstatements.SelectPostReactionStmt, []interface{}{userId, postId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("error selecting reaction: %w", err)
	} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, err
	}

	reaction, ok := reactionData[0].(*dbmodels.Reaction)
	if !ok {
		return nil, fmt.Errorf("could not assert reaction type")
	}

	return reaction, nil
}
