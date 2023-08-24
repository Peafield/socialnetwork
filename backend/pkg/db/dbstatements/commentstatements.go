package dbstatements

import (
	"database/sql"
	"fmt"
)

func initCommentDBStatements(db *sql.DB) error {
	var err error

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

	return nil
}
