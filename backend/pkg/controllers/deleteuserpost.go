package controllers

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
	column := "post_id"
	err := crud.DeleteFromDatabase(db, "Posts", column, postId)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %w", err)
	}
	return nil
}
