package dbstatements

import (
	"database/sql"
	"fmt"
)

func initCommentDBStatements(db *sql.DB) error {
	var err error

	/*INESRT*/
	InsertCommentsStmt, err = db.Prepare(`
	INSERT INTO Comments (
		comment_id,
		user_id,
		post_id,
		content,
		image_path
	) VALUES (
		?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert comments statement: %w", err)
	}

	/*SELECT*/
	SelectPostCommentsStmt, err = db.Prepare(`
	SELECT * FROM Comments
	WHERE post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select post comments statement: %w", err)
	}

	/*UPDATE*/
	UpdateCommentContent, err = db.Prepare(`
	UPDATE Comments SET content = ?
	WHERE comment_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update comment content statement: %w", err)
	}
	UpdateCommentImagePath, err = db.Prepare(`
	UPDATE Comments SET image_path = ?
	WHERE comment_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update comment image statement: %w", err)
	}

	/*DELETE*/
	DeleteUserComment, err = db.Prepare(`
	DELETE FROM Comments
	WHERE comment_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete user comment statement: %w", err)
	}

	return nil
}
