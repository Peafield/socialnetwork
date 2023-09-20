package chatcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
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

	chatMessagesData, err := crud.SelectFromDatabase(db, "Chats_Messages", dbstatements.SelectChatMessagesByChatIdStmt, []interface{}{chat.ChatId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("could not select chat messages: %w", err)
	} else if err != nil {
		return nil, err
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

func SelectGroupChatMessages(db *sql.DB, userId string, groupId string) (*dbmodels.ChatMessages, error) {
	chat, err := SelectGroupChat(db, userId, groupId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("could not select chat: %w", err)
	} else if err != nil {
		return nil, err
	}

	chatMessagesData, err := crud.SelectFromDatabase(db, "Chats_Messages", dbstatements.SelectChatMessagesByChatIdStmt, []interface{}{chat.ChatId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("could not select chat messages: %w", err)
	} else if err != nil {
		return nil, err
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
