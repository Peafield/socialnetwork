package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertPost(db *sql.DB, userId string, postData map[string]interface{}) error {
	args := []interface{}{}

	postId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create post id: %w", err)
	}
	args = append(args, postId)

	groupIdData, ok := postData["group_id"].(string)
	if ok {
		args = append(args, groupIdData)
	}

	args = append(args, userId)

	postTitle, ok := postData["title"].(string)
	if !ok {
		return fmt.Errorf("title data is not a string")
	}
	args = append(args, postTitle)

	imgPathData, ok := postData["image_path"].(string)
	if ok {
		args = append(args, imgPathData)
	}

	contentData, ok := postData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args = append(args, contentData)

	privacyLevelData, ok := postData["privacy_level"].(int)
	if !ok {
		return fmt.Errorf("privacy level data is not a int")
	}
	args = append(args, privacyLevelData)

	err = crud.InteractWithDatabase(db, dbstatements.InsertPostStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}
