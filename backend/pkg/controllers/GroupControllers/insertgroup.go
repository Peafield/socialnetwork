package groupcontrollers

import (
	"database/sql"
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertGroup(db *sql.DB, userId string, newGroupData map[string]interface{}) error {
	args := make([]interface{}, 4)

	groupId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create group id: %w", err)
	}
	args[0] = groupId

	title, ok := newGroupData["title"].(string)
	if !ok {
		return fmt.Errorf("post id is not a string")
	}
	args[1] = title

	description, ok := newGroupData["description"].(string)
	if !ok {
		return fmt.Errorf("content data is not a string")
	}
	args[2] = description

	args[3] = userId

	err = crud.InteractWithDatabase(db, dbstatements.InsertGroupsStmt, args)
	if err != nil {
		return fmt.Errorf("could not insert group: %w", err)
	}

	err = InsertGroupCreatorMember(db, userId, groupId)
	if err != nil {
		return fmt.Errorf("could not insert initial group member: %w", err)
	}

	groupChatData := make(map[string]interface{})
	groupChatData["group_id"] = groupId

	err = chatcontrollers.InsertChat(db, userId, groupChatData)
	if err != nil {
		return fmt.Errorf("could not insert initial group chat: %w", err)
	}

	return nil
}
