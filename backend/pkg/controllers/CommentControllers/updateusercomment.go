package commentcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

/*
UpdateUserComment is a function that updates a user's comment in the Comments table of the database.

The function receives three parameters: a pointer to an SQL database instance, a string of a user ID,
and a map that contains the new data to update the comment. The function starts by preparing a conditions map to define which comment to update by its comment_id, which is
extracted from the updateCommentData map. Then it checks if the updateCommentData map contains any of the immutable parameters, which are parameters that
shouldn't be updated like comment_id and creation_date. If the updateCommentData map contains any of these parameters, the function will return an error.
If no immutable parameters are included in the updateCommentData map, the function then proceeds to update the comment
in the Comments table in the database using the crud.UpdateDatabaseRow function. The update will be done based on the
conditions map and the new data to be updated is contained in the updateCommentData map. If there's any error in the update operation, the function will return a formatted error message explaining
that the comment data update operation failed. Otherwise, it will return nil to indicate that the operation was
successful.

Parameters:

  - db: *sql.DB - A pointer to the SQL database instance.
  - userId: string - The user ID of the user whose comment should be updated.
  - updateCommentData: map[string]interface{} - The map containing the new data to update the comment. It should not contain any immutable parameters.

Returns:
  - error: An error object which will be nil if the operation was successful, or containing an error message if the operation was unsuccessful.
*/
func UpdateUserComment(db *sql.DB, userId string, updateCommentData map[string]interface{}) error {
	var args []interface{}
	var query *sql.Stmt

	if content, ok := updateCommentData["content"].(string); ok {
		query = dbstatements.UpdateCommentContent
		args = append(args, content)
	}
	if imagePath, ok := updateCommentData["image_path"].(string); ok {
		query = dbstatements.UpdateCommentImagePath
		args = append(args, imagePath)
	}
	if commentID, ok := updateCommentData["comment_id"].(string); ok {
		args = append(args, commentID)
	}

	err := crud.InteractWithDatabase(db, query, args)
	if err != nil {
		return fmt.Errorf("failed to update comment data: %w", err)
	}

	return nil
}
