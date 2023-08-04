package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
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
	conditions := make(map[string]interface{})
	conditions["comment_id"] = updateCommentData["comment_id"].(string)

	immutableParameters := []string{"comment_id", "user_id", "post_id", "creation_date"}

	dataContainsImmutableParameter := helpers.MapKeyContains(updateCommentData, immutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update comment immutable parameter")
	}

	err := crud.UpdateDatabaseRow(db, "Comments", conditions, updateCommentData)
	if err != nil {
		return fmt.Errorf("failed to update comment data: %w", err)
	}

	postUpdateCommentNumConditions := make(map[string]interface{})
	postUpdateCommentNumConditions["post_id"] = updateCommentData["post_id"].(string)

	postUpdateCommentNumData := make(map[string]interface{})
	postUpdateCommentNumData["num_of_comments"] = "num_of_comments + 1"

	err = crud.UpdateDatabaseRow(db, "Posts", postUpdateCommentNumConditions, postUpdateCommentNumData)
	if err != nil {
		return fmt.Errorf("failed to update post's number of comments data: %w", err)
	}
	return nil
}
