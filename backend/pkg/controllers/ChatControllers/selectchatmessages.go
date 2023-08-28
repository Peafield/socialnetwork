package chatcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectChatMessages(db *sql.DB, userId string, receipientId string) (*dbmodels.ChatMessages, error) {
	chat, err := SelectChat(db, userId, receipientId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("could not select chat: %w", err)
	} else if err != nil {
		return nil, err
	}

	query := `SELECT * FROM Chats_Messages
	WHERE chat_id = ?`

	chatMessagesData, err := crud.SelectFromDatabase(db, "Chats_Messages", query, []interface{}{chat.ChatId})
	if err != nil {
		return nil, fmt.Errorf("error selecting chat messages: %w", err)
	}

	chatMessages := &dbmodels.ChatMessages{}
	for _, v := range chatMessagesData {
		if chatMessage, ok := v.(*dbmodels.ChatMessage); ok {
			chatMessages.ChatMessages = append(chatMessages.ChatMessages, *chatMessage)
		} else {
			return nil, fmt.Errorf("failed to assert chat message data")
		}
	}

	return chatMessages, nil
}
