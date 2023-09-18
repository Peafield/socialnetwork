package commentcontrollers

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
SelectPostComments selects all the comments related to a post.

The function first creates a query to select all comments for a post matching the post's id.
It then creates an interface slice of values to be exectuted along with the query to select the comments
from the database. It then appends these comments to a slice of comments and returns them along with an error
if one exists.

Parameters:
  - db(*sql.DB): a open database
  - postId(string): a post id

Returns:
  - comments: a slice of comment
  - error: if the database fails to return the comments or the comment data fails to be asserted.
*/
func SelectPostComments(db *sql.DB, postId string) (*dbmodels.Comments, error) {
	values := []interface{}{
		postId,
	}

	commentsData, err := crud.SelectFromDatabase(db, "Comments", dbstatements.SelectPostCommentsStmt, values)
	if err != nil && !errors.Is(err, errorhandling.ErrNoRowsAffected) {
		return nil, fmt.Errorf("failed to select post comments from database: %w", err)
	}

	comments := &dbmodels.Comments{}
	for _, v := range commentsData {
		if comment, ok := v.(*dbmodels.Comment); ok {
			commentData := &dbmodels.CommentData{}
			commentData.CommentInfo = *comment
			commentData.CommentPicture, err = os.ReadFile(comment.ImagePath)
			comments.Comments = append(comments.Comments, *commentData)
		} else {
			return nil, fmt.Errorf("failed to assert comment data")
		}
	}

	return comments, nil
}
