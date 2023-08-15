package commentcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
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
func DeleteUserComment(db *sql.DB, userId string, commentId string) error {
	args := []interface{}{}
	args = append(args, commentId)

	query := fmt.Sprintf("DELETE FROM Comments WHERE comment_id = ?")
	deleteUserCommentStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare delete User Comment statement: %w", err)
	}
	defer deleteUserCommentStatment.Close()

	//delete
	err = crud.InteractWithDatabase(db, deleteUserCommentStatment, args)
	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}
	return nil
}
