package postcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectGroupPosts(db *sql.DB, groupId string) (*dbmodels.Posts, error) {
	postsData, err := crud.SelectFromDatabase(db, "Posts", dbstatements.SelectGroupPostsStmt, []interface{}{groupId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select group posts from database: %w", err)
	}

	posts := &dbmodels.Posts{}
	for _, v := range postsData {
		if post, ok := v.(*dbmodels.Post); ok {
			postData := &dbmodels.PostData{}
			postData.PostInfo = *post
			postData.PostPicture, err = os.ReadFile(post.ImagePath)
			posts.Posts = append(posts.Posts, *postData)
		} else {
			return nil, fmt.Errorf("failed to assert post data")
		}
	}

	return posts, nil
}
