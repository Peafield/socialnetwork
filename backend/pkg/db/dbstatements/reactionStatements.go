package dbstatements

import (
	"database/sql"
	"fmt"
)

func initReactionDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
	InsertReactionsStmt, err = db.Prepare(`
	INSERT INTO Reactions (
		user_id,
		post_id,
		comment_id,
		reaction
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert reactions statement: %w", err)
	}

	/*SELECT*/
	SelectPostReactionStmt, err = db.Prepare(`
	SELECT * FROM REACTIONS
	WHERE user_id = ? AND post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select post reaction statment: %w", err)
	}

	SelectCommentReactionStmt, err = db.Prepare(`
	SELECT * FROM REACTIONS
	WHERE user_id = ? AND comment_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select comment reaction statment: %w", err)
	}

	/*UPDATE*/
	UpdatePostReaction, err = db.Prepare(`
	UPDATE Reactions SET reaction = ?
	WHERE user_id = ? AND post_id =?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update post reaction statement: %w", err)
	}
	UpdateCommentReaction, err = db.Prepare(`
	UPDATE Reactions SET reaction = ?
	WHERE user_id = ? AND comment_id =?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update comment reaction statement: %w", err)
	}

	UpdatePostIncreaseLikeStmt, err = db.Prepare(`
		UPDATE Posts 
		SET likes = likes + 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update increase post likes statement: %w", err)
	}

	UpdatePostDecreaseLikesStmt, err = db.Prepare(`
		UPDATE Posts 
		SET likes = likes - 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update decrease post likes statement: %w", err)
	}
	UpdatePostIncreaseDislikeStmt, err = db.Prepare(`
		UPDATE Posts 
		SET dislikes = dislikes + 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update increase post dislikes statement: %w", err)
	}

	UpdatePostDecreaseDislikesStmt, err = db.Prepare(`
		UPDATE Posts 
		SET dislikes = dislikes - 1
		WHERE post_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update decrease post dislikes statement: %w", err)
	}
	UpdateCommentIncreaseLikeStmt, err = db.Prepare(`
		UPDATE Comments 
		SET likes = likes + 1
		WHERE comment_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update increase comment likes statement: %w", err)
	}

	UpdateCommentDecreaseLikesStmt, err = db.Prepare(`
		UPDATE Comments 
		SET likes = likes - 1
		WHERE comment_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update decrease comment likes statement: %w", err)
	}
	UpdateCommentIncreaseDislikeStmt, err = db.Prepare(`
		UPDATE Comments 
		SET dislikes = dislikes + 1
		WHERE comment_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update increase comment dislikes statement: %w", err)
	}

	UpdateCommentDecreaseDislikesStmt, err = db.Prepare(`
		UPDATE Comments 
		SET dislikes = dislikes - 1
		WHERE comment_id = ?
		`)
	if err != nil {
		return fmt.Errorf("failed to prepare update decrease comment dislikes statement: %w", err)
	}

	/*DELETE*/
	DeletePostReaction, err = db.Prepare(`
	DELETE FROM Reactions
	WHERE user_id = ? AND post_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delect post reaction statement: %w", err)
	}

	DeleteCommentReaction, err = db.Prepare(`
	DELETE FROM Reactions
	WHERE user_id = ? AND comment_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delect post reaction statement: %w", err)
	}

	return nil
}
