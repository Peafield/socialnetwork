package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectUserViewablePosts(db *sql.DB, userId string) (*dbmodels.Posts, error) {
	conditionQuery := `
	privacy_level = 0
	UNION
	SELECT * FROM Posts 
	WHERE creator_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Followers F ON P.creator_id = F.followee_id
	WHERE F.follower_id = ? AND P.privacy_level = 1
	UNION
	SELECT P.* FROM Posts P
	JOIN Posts_Selected_Followers PSF ON P.post_id = PSF.post_id
	WHERE PSF.allowed_follower_id = ?`

	values := []interface{}{
		userId,
		userId,
		userId,
	}

	postsData, err := crud.SelectFromDatabase(db, "Posts", conditionQuery, values)
	if err != nil {
		return nil, fmt.Errorf("failed to select user viewable posts from database: %w", err)
	}

	posts := &dbmodels.Posts{}
	for _, v := range postsData {
		if post, ok := v.(dbmodels.Post); ok {
			posts.Posts = append(posts.Posts, post)
		} else {
			return nil, fmt.Errorf("failed to assert post data")
		}
	}

	return posts, nil
}
