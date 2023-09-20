package dbstatements

import (
	"database/sql"
	"fmt"
)

func initChatDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/
	InsertChatsStmt, err = db.Prepare(`
	INSERT INTO Chats (
		chat_id,
		sender_id,
		receiver_id,
		group_id
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert chats statement: %w", err)
	}

	InsertChatsMessagesStmt, err = db.Prepare(`
	INSERT INTO Chats_Messages (
		message_id,
		chat_id,
		sender_id,
		message
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert chats messages statement: %w", err)
	}

	/*SELECT*/

	SelectGroupChatStmt, err = db.Prepare(`
	SELECT * FROM Chats
	WHERE group_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select group chat statement: %w", err)
	}

	SelectChatMessagesByChatIdStmt, err = db.Prepare(`
	SELECT * FROM Chats_Messages
	WHERE chat_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select chat messages by chat id statement: %w", err)
	}
	SelectChatBySenderAndRecieverIdStmt, err = db.Prepare(`
	SELECT * FROM Chats
	WHERE (sender_id = ? AND receiver_id = ?)
	OR (sender_id = ? AND receiver_id = ?)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select chat by sender or reciever id statement: %w", err)
	}
	SelectAllChatsByUserIdStmt, err = db.Prepare(`
	SELECT * FROM Chats
	WHERE sender_id = ?
	OR receiver_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all chats by user id statement: %w", err)
	}

	return nil
}
