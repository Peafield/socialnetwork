package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertChatMessage(db *sql.DB, userId string, insertChatMessageData map[string]interface{}) error {
	args := make([]interface{}, 4)

	messageId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	chatId, ok := insertChatMessageData["chat_id"].(string)
	if !ok {
		return fmt.Errorf("chat_id is not a string or doesn't exist")
	}

	message, ok := insertChatMessageData["message"].(string)
	if !ok {
		return fmt.Errorf("message is not a string or doesn't exist")
	}

	args[0] = messageId
	args[1] = chatId
	args[2] = userId
	args[3] = message

	err = crud.InteractWithDatabase(db, dbstatements.InsertChatsMessagesStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert chat message: %w", err)
	}

	return nil
}
