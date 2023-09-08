package websocket

import (
	"errors"
	"fmt"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
	"sort"
)

func getMessagableUsers(msg ReadMessage, c *Client) error {
	var messagableUsers []BasicUserInfo

	followees, err := followercontrollers.SelectFolloweesOfSpecificUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving followee's of users: %w", err)
	}

	chats, err := chatcontrollers.SelectAllChatsByUser(dbutils.DB, c.UserID)
	if err != nil {
		return fmt.Errorf("error retrieving chats by user: %w", err)
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
	chatToSend := createMarshalledWriteMessage("messagable_users", map[string][]BasicUserInfo{
		"messagableUsers": messagableUsers,
	})
	c.send <- chatToSend
	return nil

}

func appendFolloweesToMessagableUsers(followees *dbmodels.Followers, userId string, messagableUsers *[]BasicUserInfo) error {
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

		*messagableUsers = append(*messagableUsers, BasicUserInfo{
			UUID:            f.FolloweeId,
			Name:            fUser.UserInfo.DisplayName,
			LoggedInStatus:  fUser.UserInfo.IsLoggedIn,
			LastMessage:     lastMessage.Message,
			LastMessageTime: lastMessage.CreationDate,
		})
	}
	return nil
}

func appendChatPartnersToMessagableUsers(chats *dbmodels.Chats, userId string, messagableUsers *[]BasicUserInfo) error {
	var err error
	for _, c := range chats.Chats {
		cUser := &dbmodels.UserProfileData{}

		if userId == c.SenderId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByIDStmt, c.ReceiverId)
		} else if userId == c.ReceiverId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByIDStmt, c.SenderId)
		}

		if err != nil {
			return fmt.Errorf("error getting user: %w", err)
		}

		lastMessage, err := getLastMessage(userId, cUser.UserInfo.UserId)
		if err != nil {
			return err
		}

		if !containsUser(*messagableUsers, cUser.UserInfo.UserId) {
			*messagableUsers = append(*messagableUsers, BasicUserInfo{
				UUID:            cUser.UserInfo.UserId,
				Name:            cUser.UserInfo.DisplayName,
				LoggedInStatus:  cUser.UserInfo.IsLoggedIn,
				LastMessage:     lastMessage.Message,
				LastMessageTime: lastMessage.CreationDate,
			})
		}

	}
	return nil
}

func containsUser(s []BasicUserInfo, userId string) bool {
	for _, v := range s {
		if v.UUID == userId {
			return true
		}
	}

	return false
}
