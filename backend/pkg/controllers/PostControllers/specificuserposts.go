package postcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectSpecificUserPosts(db *sql.DB, userId string, specifcUserId string) (*dbmodels.Posts, error) {
	values := []interface{}{
		specifcUserId,
		userId,
		specifcUserId,
		userId,
		specifcUserId,
	}

	postsData, err := crud.SelectFromDatabase(db, "Posts", dbstatements.SelectSpecificUserPosts, values)
	if err != nil && errors.Is(err, errorhandling.ErrNoRowsAffected) {
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
