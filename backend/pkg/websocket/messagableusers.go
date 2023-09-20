package websocket

import (
	"errors"
	"fmt"
	"log"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
	groupcontrollers "socialnetwork/pkg/controllers/GroupControllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
	"sort"
)

func getMessagableUsers(msg ReadMessage, c *Client) error {
	var messagableUsers []ChatInfo

	followees, err := followercontrollers.SelectFolloweesOfSpecificUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving followee's of users: %w", err)
	}

	chats, err := chatcontrollers.SelectAllChatsByUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving chats by user: %w", err)
	}

	userGroups, err := groupcontrollers.GetUserGroups(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("could not get user groups: %w", err)
	}

	for _, v := range userGroups.Groups {
		chat, err := chatcontrollers.SelectGroupChat(dbutils.DB, c.UserID, v.GroupId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			return fmt.Errorf("failed to select group chat: %w", err)
		}
		chats.Chats = append(chats.Chats, *chat)
	}

	err = appendFolloweesToMessagableUsers(followees, c.UserID, &messagableUsers)
	if err != nil {
		return err
	}

	err = appendChatPartnersToMessagableUsers(chats, c.UserID, &messagableUsers)
	if err != nil {
		return err
	}

	sort.Slice(messagableUsers, func(i, j int) bool {
		return messagableUsers[i].LastMessageTime.After(messagableUsers[j].LastMessageTime)
	})

	//create write message
	chatToSend := createMarshalledWriteMessage("messagable_users", map[string][]ChatInfo{
		"messagableUsers": messagableUsers,
	})
	c.send <- chatToSend
	return nil

}

func appendFolloweesToMessagableUsers(followees *dbmodels.Followers, userId string, messagableUsers *[]ChatInfo) error {
	for _, f := range followees.Followers {

		fUser, err := usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByIDStmt, f.FolloweeId)
		if err != nil {
			return fmt.Errorf("error getting user from followee id: %w", err)
		}

		lastMessage, err := getLastMessage(userId, fUser.UserInfo.UserId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			return err
		} else if err != nil {
			lastMessage = &dbmodels.ChatMessage{}
		}

		*messagableUsers = append(*messagableUsers, ChatInfo{
			UUID:            f.FolloweeId,
			Name:            fUser.UserInfo.DisplayName,
			LoggedInStatus:  fUser.UserInfo.IsLoggedIn,
			LastMessage:     lastMessage.Message,
			LastMessageTime: lastMessage.CreationDate,
			IsGroup:         false,
		})
	}
	return nil
}

func appendChatPartnersToMessagableUsers(chats *dbmodels.Chats, userId string, messagableUsers *[]ChatInfo) error {
	var err error
	for _, c := range chats.Chats {
		if c.GroupId != "" {
			appendGroupChat(userId, c, messagableUsers)
			if err != nil {
				return err
			}
		} else {
			err = appendUserChat(userId, c, messagableUsers)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func appendUserChat(userId string, c dbmodels.Chat, messagableUsers *[]ChatInfo) error {
	var err error
	cUser := &dbmodels.UserProfileData{}

	if userId == c.SenderId {
		cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByIDStmt, c.ReceiverId)
	} else if userId == c.ReceiverId {
		cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByIDStmt, c.SenderId)
	}

	if err != nil {
		log.Println(c.GroupId)
		log.Println(c.ReceiverId)
		log.Println(c.SenderId)
		return fmt.Errorf("error getting user: %w", err)
	}

	lastMessage, err := getLastMessage(userId, cUser.UserInfo.UserId)
	if err != nil {
		return err
	}

	if !containsUser(*messagableUsers, cUser.UserInfo.UserId) {
		*messagableUsers = append(*messagableUsers, ChatInfo{
			UUID:            cUser.UserInfo.UserId,
			Name:            cUser.UserInfo.DisplayName,
			LoggedInStatus:  cUser.UserInfo.IsLoggedIn,
			LastMessage:     lastMessage.Message,
			LastMessageTime: lastMessage.CreationDate,
			IsGroup:         false,
		})
	}

	return nil
}

func appendGroupChat(userId string, c dbmodels.Chat, messagableUsers *[]ChatInfo) error {
	var err error
	cGroup := &dbmodels.Group{}

	cGroup, err = groupcontrollers.GetGroupByID(dbutils.DB, c.GroupId)
	if err != nil {
		return err
	}

	lastMessage, err := getLastGroupMessage(userId, cGroup.GroupId)
	if err != nil {
		return err
	}

	if !containsUser(*messagableUsers, cGroup.GroupId) {
		*messagableUsers = append(*messagableUsers, ChatInfo{
			UUID:            c.GroupId,
			Name:            cGroup.Title,
			LoggedInStatus:  2,
			LastMessage:     lastMessage.Message,
			LastMessageTime: lastMessage.CreationDate,
			IsGroup:         true,
		})
	}

	return nil
}

func containsUser(s []ChatInfo, userId string) bool {
	for _, v := range s {
		if v.UUID == userId {
			return true
		}
	}

	return false
}
