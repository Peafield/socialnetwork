package commentcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

/*
DeleteUserComment prepares comment data in order to delete it from the database.

Parmeters:
  - db (*sql.DB): an open connection to a sql database.
  - userId (string): CURRENTLY NOT USED
  - commentId (string): the id of the comment

Returns:
  - error: if the comment fails to be deleted from the database
*/
func DeleteUserComment(db *sql.DB, userId string, deleteCommentData map[string]interface{}) error {
	commentId, ok := deleteCommentData["comment_id"].(string)
	if !ok {
		return fmt.Errorf("comment id missing or not a string")
	}
	postId, ok := deleteCommentData["post_id"].(string)
	if !ok {
		return fmt.Errorf("post id missing or not a string")
	}

	err := crud.InteractWithDatabase(db, dbstatements.DeleteUserComment, []interface{}{commentId})
	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}

	err = crud.InteractWithDatabase(db, dbstatements.UpdatePostDecreaseNumOfComments, []interface{}{postId})
	if err != nil {
		return fmt.Errorf("failed to update num of comments for a post in database: %w", err)
	}

	return nil
}
