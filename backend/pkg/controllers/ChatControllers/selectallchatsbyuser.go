package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectAllChatsByUser(db *sql.DB, userId string) (*dbmodels.Chats, error) {
	chatsData, err := crud.SelectFromDatabase(db, "Chats", dbstatements.SelectAllChatsByUserIdStmt, []interface{}{userId, userId})
	if err != nil {
		return nil, fmt.Errorf("error selecting chats: %w", err)
	}

	chats := &dbmodels.Chats{}
	for _, v := range chatsData {
		if chat, ok := v.(*dbmodels.Chat); ok {
			chats.Chats = append(chats.Chats, *chat)
		} else {
			return nil, fmt.Errorf("failed to assert chat data")
		}
	}

	return chats, nil
}
