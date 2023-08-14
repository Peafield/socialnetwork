package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

func InsertPost(db *sql.DB, userId string, postData map[string]interface{}) error {
	args := make([]interface{}, 8)

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
	userData, err := crud.SelectFromDatabase(db, "Users", dbstatements.SelectUserByID, []interface{}{userId})
	if err != nil {
		return fmt.Errorf("failed to select user from database when finding display name")
	}
	user, ok := userData[0].(*dbmodels.User)
	if !ok {
		return fmt.Errorf("could not assert user type")
	}
	userDisplayName := user.DisplayName
	args[3] = userDisplayName

	postTitle, ok := postData["title"].(string)
	if !ok {
		return fmt.Errorf("title data is not a string")
	}
	args[4] = postTitle

	imgPathData, ok := postData["image_path"].(string)
	if ok {
		args[5] = imgPathData
	} else {
		args[5] = ""
	}

	contentData, ok := postData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args[6] = contentData

	privacyLevelData, ok := postData["privacy_level"].(int)
	if !ok {
		return fmt.Errorf("privacy level data is not a int")
	}
	args[7] = privacyLevelData

	err = crud.InteractWithDatabase(db, dbstatements.InsertPostStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}
