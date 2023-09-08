package websocket

import (
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
	"socialnetwork/pkg/db/dbutils"
)

func handleOnlineUser(msg ReadMessage, c *Client) error {
	clientsToNotify := make([]Client, 0)

	followers, err := followercontrollers.SelectFollowersOfSpecificUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving follower's of users: %w", err)
	}

	chats, err := chatcontrollers.SelectAllChatsByUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving chats by user: %w", err)
	}

	for _, follower := range followers.Followers {
		client := c.hub.GetClientByID(follower.FollowerId)
		if client != nil && !containsClient(clientsToNotify, client.UserID) {
			clientsToNotify = append(clientsToNotify, *client)
		}
	}

	for _, chat := range chats.Chats {
		if c.UserID == chat.SenderId {
			client := c.hub.GetClientByID(chat.ReceiverId)
			if client != nil && !containsClient(clientsToNotify, client.UserID) {
				clientsToNotify = append(clientsToNotify, *client)
			}
		} else if c.UserID == chat.ReceiverId {
			client := c.hub.GetClientByID(chat.SenderId)
			if client != nil && !containsClient(clientsToNotify, client.UserID) {
				clientsToNotify = append(clientsToNotify, *client)
			}
		}
	}

	//send message that user is online to all of the online clients
	for _, client := range clientsToNotify {
		chatToSend := createMarshalledWriteMessage("online_user", map[string]interface{}{
			"username": c.Username,
			"online":   msg.Info["online"],
		})
		client.send <- chatToSend
	}

	return nil
}

func containsClient(s []Client, userId string) bool {
	for _, v := range s {
		if v.UserID == userId {
			return true
		}
	}

	return false
}
