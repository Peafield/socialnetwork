package websocket

import (
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"sort"
)

func getLastMessage(userId string, receiverId string) (*dbmodels.ChatMessage, error) {
	//get chats
	chatMessages, err := chatcontrollers.SelectChatMessages(dbutils.DB, userId, receiverId)
	if err != nil {
		return nil, err
	}

	sort.Slice(chatMessages.ChatMessages, func(i, j int) bool {
		return chatMessages.ChatMessages[i].CreationDate.After(chatMessages.ChatMessages[j].CreationDate)
	})

	return &chatMessages.ChatMessages[0], nil

}
