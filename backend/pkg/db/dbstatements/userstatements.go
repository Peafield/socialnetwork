package dbstatements

import (
	"database/sql"
	"fmt"
)

func initUserDBStatements(db *sql.DB) error {
	var err error

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

	err = initUserUpdateStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare update users statement: %w", err)
	}

	initUserSelectStatements()

	return nil
}

func initUserSelectStatements() {
	SelectUserByID = `SELECT * FROM Users WHERE user_id = ?`
	SelectUserByDisplayName = `SELECT * FROM Users WHERE display_name = ?`
}

func initUserUpdateStatements(db *sql.DB) error {
	var err error

	UpdateUserAccount, err = db.Prepare(`
	UPDATE Users
	SET email = ?,
	display_name = ?,
	hashed_password = ?,
	first_name = ?,
	last_name = ?,
	date_of_birth = ?,
	avatar_path = ?,
	about_me = ?
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

	return nil
}
