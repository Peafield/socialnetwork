package websocket

import (
	"fmt"
	notificationcontrollers "socialnetwork/pkg/controllers/NotificationControllers"
	"socialnetwork/pkg/db/dbutils"
	"sort"
)

func handleNotification(msg ReadMessage, c *Client) error {
	err := notificationcontrollers.InsertNewNotification(dbutils.DB, c.UserID, msg.Info)
	if err != nil {
		return fmt.Errorf("could not insert new notification: %w", err)
	}

	//need to get notify all members of group or event

	receiverId, ok := msg.Info["receiver"].(string)
	if !ok {
		return fmt.Errorf("receiver id is not a string or doesn't exist")
	}

	clientToNotify := c.hub.GetClientByID(receiverId)
	if clientToNotify == nil {
		return nil
	}

	notifications, err := notificationcontrollers.SelectAllUserNotifications(dbutils.DB, receiverId)
	if err != nil {
		return fmt.Errorf("could not select all user notifications: %w", err)
	}

	//sort the chat messages based on time
	sort.Slice(notifications.Notifications, func(i, j int) bool {
		return notifications.Notifications[i].CreationDate.After(notifications.Notifications[j].CreationDate)
	})

	messageToSend := createMarshalledWriteMessage("notification", notifications.Notifications)
	clientToNotify.send <- messageToSend

	return nil
}

func handleOpenNotifications(msg ReadMessage, c *Client) error {
	notifications, err := notificationcontrollers.SelectAllUserNotifications(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("could not select all user notifications: %w", err)
	}

	err = notificationcontrollers.UpdateAllNotificationsReadStatus(dbutils.DB, c.UserID, notifications, 1)
	if err != nil {
		return fmt.Errorf("could not update all user notifications: %w", err)
	}

	notifications, err = notificationcontrollers.SelectAllUserNotifications(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("could not select all user notifications: %w", err)
	}

	//sort the chat messages based on time
	sort.Slice(notifications.Notifications, func(i, j int) bool {
		return notifications.Notifications[i].CreationDate.After(notifications.Notifications[j].CreationDate)
	})

	messageToSend := createMarshalledWriteMessage("open_notifications", notifications.Notifications)
	c.send <- messageToSend

	return nil
}
