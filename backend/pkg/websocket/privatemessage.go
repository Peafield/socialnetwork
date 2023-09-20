package websocket

import (
	"errors"
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	groupcontrollers "socialnetwork/pkg/controllers/GroupControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
	"sort"
)

func handlePrivateMessage(msg ReadMessage, c *Client) error {
	insertChatMessageData := make(map[string]interface{}, 0)
	chat := &dbmodels.Chat{}
	var err error

	message, ok := msg.Info["message"].(string)
	if !ok {
		return fmt.Errorf("message is not a string or doesn't exist")
	}

	receiverId, r_ok := msg.Info["receiver"].(string)
	if r_ok && receiverId != "" {
		chat, err = chatcontrollers.SelectChat(dbutils.DB, c.UserID, receiverId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			return fmt.Errorf("error selecting chat: %w", err)
		}
	}

	groupId, g_ok := msg.Info["group_id"].(string)
	if g_ok && groupId != "" {
		chat, err = chatcontrollers.SelectGroupChat(dbutils.DB, c.UserID, groupId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			return fmt.Errorf("error selecting group chat: %w", err)
		}
	}

	if !r_ok && !g_ok {
		return fmt.Errorf("could not find receiver id or group id")
	}

	insertChatMessageData["chat_id"] = chat.ChatId
	insertChatMessageData["message"] = message

	err = chatcontrollers.InsertChatMessage(dbutils.DB, c.UserID, insertChatMessageData)
	if err != nil {
		return fmt.Errorf("error inserting chat message: %w", err)
	}

	chatMessages := &dbmodels.ChatMessages{}
	if r_ok && receiverId != "" {
		chatMessages, err = chatcontrollers.SelectChatMessages(dbutils.DB, c.UserID, receiverId)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}
	} else if g_ok && groupId != "" {
		chatMessages, err = chatcontrollers.SelectGroupChatMessages(dbutils.DB, c.UserID, groupId)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}
	}

	//sort the chat messages based on time
	sort.Slice(chatMessages.ChatMessages, func(i, j int) bool {
		return chatMessages.ChatMessages[i].CreationDate.Before(chatMessages.ChatMessages[j].CreationDate)
	})

	//create write message
	messageToSend := createMarshalledWriteMessage("private_message", chatMessages.ChatMessages)
	c.send <- messageToSend

	if r_ok && receiverId != "" {
		receipientClient := c.hub.GetClientByID(receiverId)
		if receipientClient != nil {
			receipientClient.send <- messageToSend
		}
	} else if g_ok && groupId != "" {
		groupMembers, err := groupcontrollers.SelectAllGroupMembers(dbutils.DB, groupId)
		if err != nil {
			return fmt.Errorf("failed to get group members: %w", err)
		}
		for _, v := range groupMembers.GroupMembers {
			if v.MemberId != c.UserID {
				receipientClient := c.hub.GetClientByID(v.MemberId)
				if receipientClient != nil {
					receipientClient.send <- messageToSend
				}
			}
		}
	}

	return nil
}
