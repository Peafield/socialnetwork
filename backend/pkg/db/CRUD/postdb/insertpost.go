package postdb

import (
	"database/sql"
	"fmt"
)

/*
InsertPost creates a new post record in the database.

It takes as input an object that implements the DBOpener interface and a User object. The DBOpener interface
should provide methods to retrieve the driver name and the data source name required for establishing a connection
with the database. The User object contains all necessary data for a new user record.

# It takes an

Upon a successful operation, the function returns the User object that was passed as an input.

Parameters:
  - db (*sql.DB): An open database to access and interact with.
  - user (*dbmodels.User): A pointer to the User object that contains data for the new user record.

Errors:
  - Returns an error if the database fails to open, if the insert statement fails to prepare or execute, or if the function fails to retrieve the number of affected rows.
  - Returns an error if no rows were affected when attempting to insert the new user.

Example:
  - InsertUser is called when a new user is created on the application, to insert the new user's data into the database.
*/
func InsertPost(db *sql.DB, values []interface{}) error {
	statement, err := db.Prepare(`
	INSERT INTO Posts (
		post_id,
		grou_id,
		creator_id,
		title,
		image_path,
		content, 
		privacy_level,
		allowed_followers,
		likes,
		dislikes
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert post statement: %s", err)
	}
	defer statement.Close()

	result, err := statement.Exec(values...)
	if err != nil {
		return fmt.Errorf("failed to execute insert post statement: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %s", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when inserting post: %s", err)
	}
	return nil
}
