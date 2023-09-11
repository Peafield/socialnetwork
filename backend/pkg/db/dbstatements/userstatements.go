package dbstatements

import (
	"database/sql"
	"fmt"
)

func initUserDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
	InsertUserStmt, err = db.Prepare(`
	INSERT INTO Users (
		user_id,
		is_logged_in,
		email,
		display_name,
		hashed_password,
		first_name,
		last_name, 
		date_of_birth,
		avatar_path,
		about_me
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert users statement: %w", err)
	}

	/*SELECT*/
	SelectAllUsersStmt, err = db.Prepare(`
	SELECT * FROM Users
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all users statement")
	}
	SelectUserByIDStmt, err = db.Prepare(`
	SELECT * FROM 
	Users WHERE user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select user by id statement: %w", err)
	}
	SelectUserByDisplayNameStmt, err = db.Prepare(`
	SELECT * FROM 
	Users WHERE
	display_name = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select user by displayname statement: %w", err)
	}

	SelectUserByEmailAndDisplayNameAndUserIdStmt, err = db.Prepare(`
	SELECT * FROM Users 
	WHERE email = ?
	AND display_name = ?
	AND user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select user by id AND display name and email statement: %w", err)
	}

	SelectUserByIDOrDisplayNameStmt, err = db.Prepare(`
	SELECT * FROM Users
	WHERE user_id = ?
	OR display_name = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select user by id OR display name statement: %w", err)
	}

	SelectUserByEmailOrDisplayNameStmt, err = db.Prepare(`
	SELECT * FROM Users
	WHERE email = ?
	OR display_name = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select user by id OR display name statement: %w", err)
	}

	UpdateUserAccountStmt, err = db.Prepare(`
	UPDATE Users
	SET email = ?,
	display_name = ?,
	hashed_password = ?,
	first_name = ?,
	last_name = ?,
	date_of_birth = ?,
	avatar_path = ?,
	about_me = ?,
	is_private = ?
	WHERE user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update user account statement: %w", err)
	}

	UpdateUserLoggedIn, err = db.Prepare(`
	UPDATE Users
	SET is_logged_in = 1
	WHERE user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update user logged in statement: %w", err)
	}

	UpdateUserLoggedOut, err = db.Prepare(`
	UPDATE Users
	SET is_logged_in = 0
	WHERE user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update user logged out statement: %w", err)
	}

	UpdateAllUsersToSignedOut, err = db.Prepare(`
	UPDATE Users
	SET is_logged_in = 0
	WHERE is_logged_in = 1
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update all users to signed out statement: %w", err)
	}

	/*DELETE*/
	DeleteUserAccountStmt, err = db.Prepare(`
	DELETE FROM Users 
	WHERE user_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete user account statement")
	}

	return nil
}
