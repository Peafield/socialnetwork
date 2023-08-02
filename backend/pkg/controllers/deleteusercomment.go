package controllers

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
	column := "comment_id"
	err := crud.DeleteFromDatabase(db, "Posts", column, commentId)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %w", err)
	}
	return nil
}
