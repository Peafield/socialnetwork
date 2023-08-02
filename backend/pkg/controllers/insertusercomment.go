package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

func InsertComment(db *sql.DB, userId string, commentData map[string]interface{}) error {
	var comment dbmodels.Comment

	commentId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create comment id: %w", err)
	}
	comment.CommentId = commentId

	comment.UserId = userId

	postIdData, ok := commentData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	comment.PostId = postIdData

	contentData, ok := commentData["content"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	comment.Content = contentData

	imgPathData, ok := commentData["image_path"].(string)
	if ok {
		comment.ImagePath = imgPathData
	}

	values, err := helpers.StructFieldValues(comment)
	if err != nil {
		return fmt.Errorf("failed to get comment struct values: %s", err)
	}
	err = crud.InsertIntoDatabase(db, dbstatements.InsertCommentsStmt, values)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %s", err)
	}
	return nil
}
