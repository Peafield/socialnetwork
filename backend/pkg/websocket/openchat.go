package websocket

import (
	"errors"
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"sort"
)

func handleOpenChat(msg ReadMessage, c *Client) error {
	//assert msg info type
	receiverId, ok := msg.Info["receiver"].(string)
	if !ok {
		return fmt.Errorf("receiver id is not a string")
	}

	err := createChatIfNoneFound(c, receiverId)
	if err != nil {
		return fmt.Errorf("could not create chat: %w", err)
	}

	//get chats
	chatMessages, err := chatcontrollers.SelectChatMessages(dbutils.DB, c.UserID, receiverId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return fmt.Errorf("error: %v", err)
	} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
		err = sendChatID(c, receiverId)
		if err != nil {
			return fmt.Errorf("error sending chat id: %w", err)
		}
		return nil
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

func createChatIfNoneFound(c *Client, receiverId string) error {
	_, err := chatcontrollers.SelectChat(dbutils.DB, c.UserID, receiverId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return fmt.Errorf("error selecting chat: %w", err)
	} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
		err = chatcontrollers.InsertChat(dbutils.DB, c.UserID, map[string]interface{}{"receiver_id": receiverId})
		if err != nil {
			return fmt.Errorf("error inserting chat: %w", err)
		}
	}
	return nil
}

func sendChatID(c *Client, receiverId string) error {
	chat, err := chatcontrollers.SelectChat(dbutils.DB, c.UserID, receiverId)
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return fmt.Errorf("error selecting chat: %w", err)
	}
	chatToSend := createMarshalledWriteMessage("open_chat", chat.ChatId)
	c.send <- chatToSend
	return nil
}
