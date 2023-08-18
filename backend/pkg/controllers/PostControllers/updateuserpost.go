package postcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"strings"
)

/*
UpdateUserPost is a function that updates a user's post in the Posts table of the database.

The function receives three parameters: a pointer to an SQL database instance, a string of a user ID,
and a map that contains the new data to update the post. The function starts by preparing a conditions map to define which post to update by its post_id, which is
extracted from the updatePostData map. Then it checks if the updatePostData map contains any of the immutable parameters, which are parameters that
shouldn't be updated like post_id, group_id, creator_id, likes, dislikes, and creation_date. If the updatePostData
map contains any of these parameters, the function will return an error. If no immutable parameters are included in the updatePostData map, the function then proceeds to update the post
in the Posts table in the database using the crud.UpdateDatabaseRow function. The update will be done based on the
conditions map and the new data to be updated is contained in the updatePostData map. If there's any error in the update operation, the function will return a formatted error message explaining
that the post data update operation failed. Otherwise, it will return nil to indicate that the operation was
successful.

Parameters:

	-db: *sql.DB - A pointer to the SQL database instance.
	-userId: string - The user ID of the user whose post should be updated.
	- updatePostData: map[string]interface{} - The map containing the new data to update the post. It should not contain any immutable parameters.

Returns:
  - error: An error object which will be nil if the operation was successful, or containing an error message if the operation was unsuccessful.
*/
func UpdateUserPost(db *sql.DB, userId string, updatePostData map[string]interface{}) error {
	var columns []string
	var args []interface{}

	if title, ok := updatePostData["title"].(string); ok {
		columns = append(columns, "title = ?")
		args = append(args, title)
	}
	if imagePath, ok := updatePostData["image_path"].(string); ok {
		columns = append(columns, "image_path = ?")
		args = append(args, imagePath)
	}
	if content, ok := updatePostData["content"].(string); ok {
		columns = append(columns, "content = ?")
		args = append(args, content)
	}
	if numOfComments, ok := updatePostData["num_of_comments"].(int); ok {
		columns = append(columns, "num_of_comments = ?")
		args = append(args, numOfComments)
	}
	if privacyLevel, ok := updatePostData["privacy_level"].(string); ok {
		columns = append(columns, "privacy_level = ?")
		args = append(args, privacyLevel)
	}
	if postId, ok := updatePostData["post_id"].(string); ok {
		args = append(args, postId)
	}

	query := fmt.Sprintf("UPDATE Posts SET %s WHERE post_id = ?", strings.Join(columns, ", "))
	updatePostStatement, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update post statement: %w", err)
	}
	defer updatePostStatement.Close()

	err = crud.InteractWithDatabase(db, updatePostStatement, args)
	if err != nil {
		return fmt.Errorf("failed to update post data: %w", err)
	}
	return nil
}
