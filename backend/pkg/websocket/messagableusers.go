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

func getMessagableUsers(userId string) ([]BasicUserInfo, error) {
	var messagableUsers []BasicUserInfo

	followees, err := followercontrollers.SelectFolloweesOfSpecificUser(dbutils.DB, userId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving followee's of users: %w", err)
	}

	chats, err := chatcontrollers.SelectAllChatsByUser(dbutils.DB, userId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving chats by user: %w", err)
	}

	for _, f := range followees.Followers {

		fUser, err := usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, f.FolloweeId)
		if err != nil {
			return nil, fmt.Errorf("error getting user from followee id: %w", err)
		}

		lastMessage, err := getLastMessage(userId, fUser.UserInfo.UserId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			return nil, err
		} else if err != nil {
			lastMessage = &dbmodels.ChatMessage{}
		}

		messagableUsers = append(messagableUsers, BasicUserInfo{
			UUID:            f.FolloweeId,
			Name:            fUser.UserInfo.DisplayName,
			LoggedInStatus:  fUser.UserInfo.IsLoggedIn,
			LastMessage:     lastMessage.Message,
			LastMessageTime: lastMessage.CreationDate,
		})
	}

	for _, c := range chats.Chats {
		cUser := &dbmodels.UserProfileData{}

		if userId == c.SenderId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, c.ReceiverId)
		} else if userId == c.ReceiverId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, c.SenderId)
		}

		if err != nil {
			return nil, fmt.Errorf("error getting user: %w", err)
		}

		lastMessage, err := getLastMessage(userId, cUser.UserInfo.UserId)
		if err != nil {
			return nil, err
		}

		if !containsUser(messagableUsers, cUser.UserInfo.UserId) {
			messagableUsers = append(messagableUsers, BasicUserInfo{
				UUID:            cUser.UserInfo.UserId,
				Name:            cUser.UserInfo.DisplayName,
				LoggedInStatus:  cUser.UserInfo.IsLoggedIn,
				LastMessage:     lastMessage.Message,
				LastMessageTime: lastMessage.CreationDate,
			})
		}

	}

	sort.Slice(messagableUsers, func(i, j int) bool {
		return messagableUsers[i].LastMessageTime.After(messagableUsers[j].LastMessageTime)
	})

	return messagableUsers, nil
}

func containsUser(s []BasicUserInfo, userId string) bool {
	for _, v := range s {
		if v.UUID == userId {
			return true
		}
	}

	return false
}
