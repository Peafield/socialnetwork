package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertGroupChat(db *sql.DB, userId string, insertChatData map[string]interface{}) error {
	args := make([]interface{}, 2)

	chatId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("error creating group chat uuid: %w", err)
	}

	groupId, ok := insertChatData["group_id"].(string)
	if !ok {
		return fmt.Errorf("group_id is not a string or doesn't exist")
	}

	args[0] = chatId
	args[1] = groupId

	err = crud.InteractWithDatabase(db, dbstatements.InsertGroupChatStmt, args)
	if err != nil {
		return fmt.Errorf("error inserting group chat: %w", err)
	}
	return nil
}
