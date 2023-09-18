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

/*
SelectUserViewablePosts selects all user viewable posts.

The function first creates a query to select all user viewable posts match:
  - where the privacy level is public (0)
  - where the user is the creator of the post
  - where the user is a follower of the post's creator
  - where the user is one of the selected followers of the post

It then creates an interface slice of values to be exectuted along with the query to select the posts
from the database. It then appends these posts to a slice of posts and returns them along with an error
if one exists.

Parameters:
  - db (*sql.DB): a open database
  - userId (string): a user id

Returns:
  - posts: a slice of posts
  - error: if the database fails to return the posts or the post data fails to be asserted.
*/
func SelectUserViewablePosts(db *sql.DB, userId string) (*dbmodels.Posts, error) {
	values := []interface{}{
		userId,
		userId,
		userId,
		userId,
	}

	postsData, err := crud.SelectFromDatabase(db, "Posts", dbstatements.SelectUserViewablePostsStmt, values)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select user viewable posts from database: %w", err)
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
