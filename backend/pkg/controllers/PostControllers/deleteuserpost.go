package postcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
)

/*
DeleteUserPost prepares post data in order to delete it from the database.

Parmeters:
  - db (*sql.DB): an open connection to a sql database.
  - userId (string): CURRENTLY NOT USED
  - postId (string): the id of the post

Returns:
  - error: if the post fails to be deleted from the database
*/
func DeleteUserPost(db *sql.DB, userId string, postId string) error {
	args := []interface{}{}
	args = append(args, postId)

	query := fmt.Sprintf("DELETE FROM Posts WHERE post_id = ?")
	deleteUserPostStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare delete User Post statement: %w", err)
	}
	defer deleteUserPostStatment.Close()

	//delete
	err = crud.InteractWithDatabase(db, deleteUserPostStatment, args)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %w", err)
	}
	return nil
}
