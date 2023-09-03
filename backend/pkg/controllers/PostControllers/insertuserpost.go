package postcontrollers

import (
	"database/sql"
	"fmt"
	imagecontrollers "socialnetwork/pkg/controllers/ImageControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertPost(db *sql.DB, userId string, postData map[string]interface{}) error {
	args := make([]interface{}, 6)

	postId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create post id: %w", err)
	}
	args[0] = postId

	groupIdData, ok := postData["group_id"].(string)
	if ok {
		args[1] = groupIdData
	} else {
		args[1] = ""
	}

	args[2] = userId

	imgPathData, ok := postData["image_path"].(string)
	if ok {
		imgFilePath, err := imagecontrollers.DecodeImage(imgPathData)
		if err != nil {
			return fmt.Errorf("problem decoding image")
		}
		args[3] = imgFilePath
	} else {
		args[3] = ""
	}

	contentData, ok := postData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args[4] = contentData

	privacyLevelData, ok := postData["privacy_level"].(float64)
	if !ok {
		return fmt.Errorf("privacy level data is not a string")
	}

	args[5] = int(privacyLevelData)

	err = crud.InteractWithDatabase(db, dbstatements.InsertPostStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}
