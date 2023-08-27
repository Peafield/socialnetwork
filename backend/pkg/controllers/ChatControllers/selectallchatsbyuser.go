package chatcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectAllChatsByUser(db *sql.DB, userId string) (*dbmodels.Chats, error) {
	query := `SELECT * FROM Chats
	WHERE sender_id = ?
	OR receiver_id = ?`

	chatsData, err := crud.SelectFromDatabase(db, "Chats", query, []interface{}{userId, userId})
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
