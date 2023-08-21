package dbstatements

import (
	"database/sql"
	"fmt"
)

func initPostDBStatements(db *sql.DB) error {
	var err error

	InsertPostStmt, err = db.Prepare(`
	INSERT INTO Posts (
		post_id,
		group_id,
		creator_id,
		image_path,
		content,
		privacy_level
	) VALUES (
		?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert posts statement: %w", err)
	}

	err = initPostUpdateStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare update users statement: %w", err)
	}

	initPostSelectStatements()

	return nil
}

func initPostUpdateStatements(db *sql.DB) error {
	var err error

	UpdatePostNumOfComments, err = db.Prepare(`
		UPDATE Posts 
		SET num_of_comments = num_of_comments + 1 
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update number of post comments statement: %w", err)
	}

	return nil
}

func initPostSelectStatements() {
	SelectUserViewablePosts = `
	SELECT * FROM Posts
	WHERE privacy_level = 0
	UNION
	SELECT * FROM Posts 
	WHERE creator_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Followers F ON P.creator_id = F.followee_id
	WHERE F.follower_id = ? AND P.privacy_level = 1
	UNION
	SELECT P.* FROM Posts P
	JOIN Posts_Selected_Followers PSF ON P.post_id = PSF.post_id
	WHERE PSF.allowed_follower_id = ?`

	SpecificUserPosts = `
	SELECT * FROM Posts 
	WHERE creator_id = ? AND privacy_level = 0
	UNION
	SELECT P.* FROM Posts P
	JOIN Followers F ON P.creator_id = F.followee_id
	WHERE F.follower_id = ? AND P.privacy_level = 1 AND F.followee_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Posts_Selected_Followers PSF ON P.post_id = PSF.post_id
	WHERE PSF.allowed_follower_id = ? AND P.creator_id = ?`
}
