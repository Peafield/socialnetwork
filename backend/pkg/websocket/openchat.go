package websocket

import (
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	"socialnetwork/pkg/db/dbutils"
	"sort"
)

func handleOpenChat(msg ReadMessage, c *Client) error {
	//assert msg info type
	receiverId, ok := msg.Info["receiver"].(string)
	if !ok {
		return fmt.Errorf("receiver id is not a string")
	}

	//get chats
	chatMessages, err := chatcontrollers.SelectChatMessages(dbutils.DB, c.UserID, receiverId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	//sort the chat messages based on time
	sort.Slice(chatMessages.ChatMessages, func(i, j int) bool {
		return chatMessages.ChatMessages[i].CreationDate.Before(chatMessages.ChatMessages[j].CreationDate)
	})

	//create write message
	chatToSend := createMarshalledWriteMessage("open_chat", chatMessages.ChatMessages)
	c.send <- chatToSend
	return nil
}
