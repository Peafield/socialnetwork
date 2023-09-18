package commentcontrollers

import (
	"database/sql"
	"fmt"
	imagecontrollers "socialnetwork/pkg/controllers/ImageControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertComment(db *sql.DB, userId string, commentData map[string]interface{}) error {
	args := make([]interface{}, 5)

	commentId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create comment id: %w", err)
	}
	args[0] = commentId

	args[1] = userId

	postIdData, ok := commentData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	args[2] = postIdData

	contentData, ok := commentData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args[3] = contentData

	imgPathData, ok := commentData["image"].(string)
	if ok && imgPathData != "" {
		imgFilePath, err := imagecontrollers.DecodeImage(imgPathData)
		if err != nil {
			return fmt.Errorf("problem decoding image")
		}
		args[4] = imgFilePath
	} else {
		args[4] = ""
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertCommentsStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert comment into database: %s", err)
	}

	err = crud.InteractWithDatabase(db, dbstatements.UpdatePostIncreaseNumOfComments, []interface{}{postIdData})
	if err != nil {
		return fmt.Errorf("failed to update post comment number: %s", err)
	}

	return nil
}
