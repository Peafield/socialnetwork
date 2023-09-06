package dbstatements

import (
	"database/sql"
	"fmt"
)

func initFollowerDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
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

	/*SELECT*/
	SelectFollowerInfoStmt, err = db.Prepare(`
	SELECT * FROM Followers
	WHERE follower_id = ?
	AND followee_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select follower info statement: %w", err)
	}

	SelectFolloweesOfUserStmt, err = db.Prepare(`
	SELECT * FROM Followers
	WHERE follower_id = ? AND following_status = 1
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select followees of user statement: %w", err)
	}

	SelectFollowersOfUserStmt, err = db.Prepare(`
	SELECT * FROM Followers
	WHERE followee_id = ? AND following_status = 1
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select followers of user statement: %w", err)
	}

	/*UPDATE*/
	UpdateFollowingStatusStmt, err = db.Prepare(`
	UPDATE Followers
	SET following_status = ?, request_pending = 0
	WHERE follower_id = ? AND followee_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update following status statement: %w", err)
	}

	UpdateRequestPendingStmt, err = db.Prepare(`
	UPDATE Followers SET request_pending = ?
	WHERE followee_id = ? AND Follower_id =?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update request pending statement: %w", err)
	}

	/*DELETE*/
	DeleteFollowerStmt, err = db.Prepare(`
	DELETE FROM Followers WHERE followee_id = ? AND follower_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete follower info statement: %w", err)
	}

	return nil
}
