package websocket

import (
	"errors"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
	"sort"
)

func getLastMessage(userId string, receiverId string) (*dbmodels.ChatMessage, error) {
	//get chats
	chatMessages, err := chatcontrollers.SelectChatMessages(dbutils.DB, userId, receiverId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, err
	} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
		return &dbmodels.ChatMessage{}, nil
	}

	sort.Slice(chatMessages.ChatMessages, func(i, j int) bool {
		return chatMessages.ChatMessages[i].CreationDate.After(chatMessages.ChatMessages[j].CreationDate)
	})

	return &chatMessages.ChatMessages[0], nil

}

func getLastGroupMessage(userId string, groupId string) (*dbmodels.ChatMessage, error) {
	//get chats
	chatMessages, err := chatcontrollers.SelectGroupChatMessages(dbutils.DB, userId, groupId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, err
	} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
		return &dbmodels.ChatMessage{}, nil
	}

	sort.Slice(chatMessages.ChatMessages, func(i, j int) bool {
		return chatMessages.ChatMessages[i].CreationDate.After(chatMessages.ChatMessages[j].CreationDate)
	})

	return &chatMessages.ChatMessages[0], nil

}
