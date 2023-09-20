package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertChat(db *sql.DB, userId string, insertChatData map[string]interface{}) error {
	args := make([]interface{}, 4)

	chatId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("error creating chat uuid: %w", err)
	}
	args[0] = chatId

	receiverId, ok := insertChatData["receiver_id"].(string)
	if ok {
		args[1] = userId
		args[2] = receiverId
	} else {
		args[1] = ""
		args[2] = ""
	}

	groupId, ok := insertChatData["group_id"].(string)
	if ok {
		args[3] = groupId
	} else {
		args[3] = ""
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertChatsStmt, args)
	if err != nil {
		return fmt.Errorf("error inserting chat: %w", err)
	}
	return nil
}
