package postcontrollers

import (
	"database/sql"
	"fmt"
	imagecontrollers "socialnetwork/pkg/controllers/ImageControllers"
	postselectedfollowercontrollers "socialnetwork/pkg/controllers/PostSelectedFollowerControllers"
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
	if ok && imgPathData != "" {
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
		return fmt.Errorf("privacy level data is not a float64")
	}

	args[5] = int(privacyLevelData)

	if int(privacyLevelData) == 2 {
		err := insertPostSelectedFollowers(db, postData, postId, userId)
		if err != nil {
			return err
		}
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertPostStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}

func insertPostSelectedFollowers(db *sql.DB, postData map[string]interface{}, postId string, userId string) error {
	selectedFollowers, ok := postData["selected_profiles"].([]interface{})
	if !ok {
		return fmt.Errorf("no selected profiles selected")
	}
	for _, v := range selectedFollowers {
		psf := make(map[string]interface{})
		psf["post_id"] = postId
		selectedFollower, ok := v.(string)
		if !ok {
			return fmt.Errorf("selected profile not a string or doesnt exist")
		}
		psf["allowed_follower_id"] = selectedFollower
		postselectedfollowercontrollers.NewPostSelectedFollower(db, userId, psf)
	}

	return nil
}
