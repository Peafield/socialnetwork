package dbstatements

import (
	"database/sql"
	"fmt"
)

func initNotificationStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
	InsertNotificationsStmt, err = db.Prepare(`
	INSERT INTO Notifications (
		notification_id,
		sender_id,
		receiver_id,
		group_id,
		post_id,
		event_id,
		comment_id,
		chat_id,
		reaction_type
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert notifications statement: %w", err)
	}

	/*SELECT*/
	SelectAllUserNotifications, err = db.Prepare(`
	SELECT * FROM Notifications
	WHERE receiver_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all user notifications statement: %w", err)
	}

	/*UPDATE*/
	UpdateAllUserNotifications, err = db.Prepare(`
	UPDATE Notifications
	SET read_status = ?
	WHERE notification_id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare update all user notifications statement: %w", err)
	}

	/*DELETE*/
	DeleteNotificationStmt, err = db.Prepare(`
	DELETE FROM Notifications
	WHERE notification_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete notification statement: %w", err)
	}

	return nil
}
