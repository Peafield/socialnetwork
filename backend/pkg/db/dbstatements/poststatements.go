package dbstatements

import (
	"database/sql"
	"fmt"
)

func initPostDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
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

	InsertPostsSelectedFollowerStmt, err = db.Prepare(`
	INSERT INTO Posts_Selected_Followers  (
		post_id,
		allowed_follower_id
	) VALUES (
		?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert post follower statement: %w", err)
	}

	/*SELECT*/
	SelectAllPostsStmt, err = db.Prepare(`
	SELECT * FROM Posts
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all post id statment: %w", err)
	}

	SelectUserViewablePostsStmt, err = db.Prepare(`
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
	WHERE PSF.allowed_follower_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Groups_Members GM ON P.group_id = GM.group_id
	WHERE P.privacy_level = 3 AND GM.member_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare user viewable posts statement: %w", err)
	}

	SelectSpecificUserPostsStmt, err = db.Prepare(`
	SELECT * FROM Posts 
	WHERE (creator_id = ? AND privacy_level = 0) OR (creator_id = ?)
	UNION
	SELECT P.* FROM Posts P
	JOIN Followers F ON P.creator_id = F.followee_id
	WHERE F.follower_id = ? AND P.privacy_level = 1 AND F.followee_id = ?
	UNION
	SELECT P.* FROM Posts P
	JOIN Posts_Selected_Followers PSF ON P.post_id = PSF.post_id
	WHERE PSF.allowed_follower_id = ? AND P.creator_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare user specific posts statement: %w", err)
	}

	SelectGroupPostsStmt, err = db.Prepare(`
	SELECT * FROM Posts
	WHERE privacy_level = 3 AND group_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare group posts statement: %w", err)
	}

	/*UPDATE*/
	UpdatePostImagePathStmt, err = db.Prepare(`
	UPDATE Posts
	SET image_path = ?
	WHERE post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update post image path statement: %w", err)
	}
	UpdatePostContentStmt, err = db.Prepare(`
	UPDATE Posts
	SET content = ?
	WHERE post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update post content statement: %w", err)
	}
	UpdatePostPrivacyLevelStmt, err = db.Prepare(`
	UPDATE Posts
	SET privacy_level = ?
	WHERE post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update post privacy leel statement: %w", err)
	}

	UpdatePostIncreaseNumOfComments, err = db.Prepare(`
		UPDATE Posts 
		SET num_of_comments = num_of_comments + 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update increase number of post comments statement: %w", err)
	}

	UpdatePostDecreaseNumOfComments, err = db.Prepare(`
		UPDATE Posts 
		SET num_of_comments = num_of_comments - 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update decrease number of post comments statement: %w", err)
	}
	/*DELETE*/
	DeletePostStmt, err = db.Prepare(`
	DELETE FROM Posts
	WHERE post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete User Post statement: %w", err)
	}

	DeletePostsSelectedFollowerStmt, err = db.Prepare(`
	DELETE FROM Posts_Selected_Followers 
	WHERE post_id = ? AND allowed_follower_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete post selected follower statement: %w", err)
	}

	return nil
}
