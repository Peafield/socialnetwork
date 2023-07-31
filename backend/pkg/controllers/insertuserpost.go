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
	var post dbmodels.Post

	postId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create post id: %w", err)
	}
	post.PostId = postId

	groupIdData, ok := postData["group_id"].(string)
	if ok {
		post.GroupId = groupIdData
	}

	post.CreatorId = userId

	imgPathData, ok := postData["image_path"].(string)
	if ok {
		post.ImagePath = imgPathData
	}

	contentData, ok := postData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	post.Content = contentData

	privacyLevelData, ok := postData["privacy_level"].(int)
	if !ok {
		return fmt.Errorf("privacy level dat is not a int")
	}
	post.PrivacyLevel = privacyLevelData

	allowedFollowersData, ok := postData["allowed_followers"].(string)
	if !ok {
		return fmt.Errorf("allowed followers data is not a string")
	}
	post.AllowedFollowers = allowedFollowersData

	values, err := helpers.StructFieldValues(post)
	if err != nil {
		return fmt.Errorf("failed to get post struct values: %s", err)
	}
	err = crud.InsertIntoDatabase(db, dbstatements.InsertPostStmt, values)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}