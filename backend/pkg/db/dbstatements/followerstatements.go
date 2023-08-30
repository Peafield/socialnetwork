package dbstatements

import (
	"database/sql"
	"fmt"
)

func initFollowerDBStatements(db *sql.DB) error {
	var err error

	InsertFollowersStmt, err = db.Prepare(`
	INSERT INTO Followers (
		follower_id,
		followee_id,
		following_status,
		request_pending
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert followers statement: %w", err)
	}

	err = initFollowerUpdateStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare update followers statement: %w", err)
	}

	initFollowerSelectStatements()

	return nil
}

func initFollowerUpdateStatements(db *sql.DB) error {
	var err error

	UpdateFollowerStatus, err = db.Prepare(`
	UPDATE Followers
	SET following_status = ?, request_pending = 0
	WHERE follower_id = ? AND followee_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update follower status statement: %w", err)
	}

	return nil
}

func initFollowerSelectStatements() {
	SelectFollowersOfUser = `
	SELECT * FROM Followers
	WHERE followee_id = ? AND following_status = 1
	`

	SelectFolloweesOfUser = `
	SELECT * FROM Followers
	WHERE follower_id = ? AND following_status = 1`

	SelectFollower = `SELECT * FROM Followers
	WHERE follower_id = ?
	AND followee_id = ?`
}
