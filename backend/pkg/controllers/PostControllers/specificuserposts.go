package postcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectSpecificUserPosts(db *sql.DB, userId string, specifcUserId string) (*dbmodels.Posts, error) {
	query := `
	SELECT * FROM Posts 
	WHERE creator_id = ? AND privacy_level = 0
	UNION
	SELECT P.* FROM Posts P
	JOIN Followers F ON P.creator_id = F.followee_id
	WHERE F.follower_id = ? AND P.privacy_level = 1 AND F.followee_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Posts_Selected_Followers PSF ON P.post_id = PSF.post_id
	WHERE PSF.allowed_follower_id = ? AND P.creator_id = ?`

	values := []interface{}{
		specifcUserId,
		userId,
		specifcUserId,
		userId,
		specifcUserId,
	}

	postsData, err := crud.SelectFromDatabase(db, "Posts", query, values)
	if err != nil {
		return nil, fmt.Errorf("failed to select user viewable posts from database: %w", err)
	}

	posts := &dbmodels.Posts{}
	for _, v := range postsData {
		if post, ok := v.(*dbmodels.Post); ok {
			posts.Posts = append(posts.Posts, *post)
		} else {
			return nil, fmt.Errorf("failed to assert post data")
		}
	}

	return posts, nil
}
