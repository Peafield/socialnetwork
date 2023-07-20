package userdb

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/models/dbmodels"
)

/*
InsertUser creates a new user record in the database.

It takes as input an object that implements the DBOpener interface and a User object. The DBOpener interface
should provide methods to retrieve the driver name and the data source name required for establishing a connection
with the database. The User object contains all necessary data for a new user record.

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
func InsertUser(db *sql.DB, user *dbmodels.User) error {
	statement, err := db.Prepare(`
	INSERT INTO Users (
		user_id,
		isLoggedIn,
		email,
		hashed_password,
		first_name,
		last_name, 
		date_of_birth,
		avatar_path,
		display_name,
		about_me
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert user statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.UserId, user.IsLoggedIn, user.Email, user.HashedPassword, user.FirstName, user.LastName, user.DOB, user.AvatarPath, user.DisplayName, user.AboutMe)
	if err != nil {
		return fmt.Errorf("failed to execute insert user statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when inserting user: %w", err)
	}

	return nil
}
