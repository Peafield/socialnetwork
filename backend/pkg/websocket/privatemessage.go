package websocket

import (
	"errors"
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
)

func handlePrivateMessage(msg ReadMessage, c *Client) error {
	insertChatMessageData := make(map[string]interface{}, 0)

	receiverId, ok := msg.Info["receiver"].(string)
	if !ok {
		return fmt.Errorf("receiver id is not a string or doesn't exist")
	}

	message, ok := msg.Info["message"].(string)
	if !ok {
		return fmt.Errorf("message is not a string or doesn't exist")
	}

	chat, err := chatcontrollers.SelectChat(dbutils.DB, c.UserID, receiverId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return fmt.Errorf("error selecting chat: %w", err)
	} else {
		err = chatcontrollers.InsertChat(dbutils.DB, c.UserID, map[string]interface{}{"receiver_id": receiverId})
		if err != nil {
			return fmt.Errorf("error inserting chat: %w", err)
		}

		chat, err = chatcontrollers.SelectChat(dbutils.DB, c.UserID, receiverId)
		if err != nil {
			return fmt.Errorf("error selecting chat: %w", err)
		}
	}

	insertChatMessageData["chat_id"] = chat.ChatId
	insertChatMessageData["message"] = message

	err = chatcontrollers.InsertChatMessage(dbutils.DB, c.UserID, insertChatMessageData)
	if err != nil {
		return fmt.Errorf("error inserting chat message: %w", err)
	}

	chatMessages, err := chatcontrollers.SelectChatMessages(dbutils.DB, c.UserID, receiverId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	receipientClient := c.hub.GetClientByID(receiverId)

	//create write message
	messageToSend := createMarshalledWriteMessage("private_message", chatMessages.ChatMessages)
	c.send <- messageToSend
	if receipientClient != nil {
		receipientClient.send <- messageToSend
	}

	return nil
}
