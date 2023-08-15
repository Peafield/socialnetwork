package commentcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertComment(db *sql.DB, userId string, commentData map[string]interface{}) error {
	args := []interface{}{}

	commentId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create comment id: %w", err)
	}
	args = append(args, commentId)

	args = append(args, userId)

	postIdData, ok := commentData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	args = append(args, postIdData)

	contentData, ok := commentData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args = append(args, contentData)

	imgPathData, ok := commentData["image_path"].(string)
	if ok {
		args = append(args, imgPathData)
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertCommentsStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}
