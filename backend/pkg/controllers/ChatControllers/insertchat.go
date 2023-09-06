package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertChat(db *sql.DB, userId string, insertChatData map[string]interface{}) error {
	args := make([]interface{}, 3)

	chatId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("error creating chat uuid: %w", err)
	}

	receiverId, ok := insertChatData["receiver_id"].(string)
	if !ok {
		return fmt.Errorf("receiver_id is not a string or doesn't exist")
	}

	args[0] = chatId
	args[1] = userId
	args[2] = receiverId

	err = crud.InteractWithDatabase(db, dbstatements.InsertChatsStmt, args)
	if err != nil {
		return fmt.Errorf("error inserting chat: %w", err)
	}
	return nil
}
