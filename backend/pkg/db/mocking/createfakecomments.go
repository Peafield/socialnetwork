package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	commentcontrollers "socialnetwork/pkg/controllers/CommentControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/models/dbmodels"
)

func CreateFakeComments(db *sql.DB) error {
	commentAmount := 60
	commentData := map[string]interface{}{}

	userIds, err := GetAllUserIDs(db)
	if err != nil {
		return fmt.Errorf("failed to get all user id's when faking comments: %w", err)
	}

	postIds, err := GetAllPostIDs(db)
	if err != nil {
		return fmt.Errorf("failed to get all post id's when faking comments: %w", err)
	}

	numOfPosts := len(postIds)
	numOfUsers := len(userIds)

	for i := 1; i <= commentAmount; i++ {
		content := ContentGenerator()
		commentData["content"] = content[:100]

		commentData["post_id"] = postIds[rand.Intn(numOfPosts)]

		err := commentcontrollers.InsertComment(db, userIds[rand.Intn(numOfUsers)], commentData)
		if err != nil {
			return fmt.Errorf("failed to insert mock comment data: %w", err)
		}
	}
	return nil
}

func GetAllUserIDs(db *sql.DB) ([]string, error) {
	userIDData, err := crud.SelectFromDatabase(db, "Users", dbstatements.SelectAllUsersStmt, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get userId's from db: %w", err)
	}

	var userIds []string
	for _, v := range userIDData {
		if user, ok := v.(*dbmodels.User); ok {
			userIds = append(userIds, user.UserId)
		} else {
			return nil, fmt.Errorf("failed to assert user id data")
		}
	}

	return userIds, nil
}

func GetAllPostIDs(db *sql.DB) ([]string, error) {
	postIDData, err := crud.SelectFromDatabase(db, "Posts", dbstatements.SelectAllPostsStmt, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get postId's from db: %w", err)
	}

	var postIds []string
	for _, v := range postIDData {
		if post, ok := v.(*dbmodels.Post); ok {
			postIds = append(postIds, post.PostId)
		} else {
			return nil, fmt.Errorf("failed to assert post id data")
		}
	}

	return postIds, nil
}
