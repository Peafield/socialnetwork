package controllers

import (
	"database/sql"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectUserViewablePosts(db *sql.DB, userId string) (*dbmodels.Posts, error) {
	var posts dbmodels.Posts

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

	values := []string{userId, userId, userId}

	return &posts, nil
}
