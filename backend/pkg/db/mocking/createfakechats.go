package db

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
)

func CreateFakeChats(db *sql.DB) error {
	chatData := map[string]interface{}{}
	chatMessageData := map[string]interface{}{}

	userIds, err := GetAllUserIDs(db)
	if err != nil {
		return fmt.Errorf("failed to get all user id's when faking chats: %w", err)
	}

	numOfUsers := len(userIds)

	for i := 0; i < numOfUsers; i++ {
		currentUserId := i
		followees, err := followercontrollers.SelectFolloweesOfSpecificUser(dbutils.DB, userIds[currentUserId])
		if err != nil {
			return fmt.Errorf("failed to select followee's when mocking")
		}

		for j := 0; len(followees.Followers) > 1; j++ {
			numOfFollowees := len(followees.Followers)
			randomUser := rand.Intn(numOfFollowees - 1)

			chatData["receiver_id"] = followees.Followers[randomUser].FolloweeId

			chat, err := chatcontrollers.SelectChat(dbutils.DB, userIds[currentUserId], followees.Followers[randomUser].FolloweeId)
			if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
				return fmt.Errorf("error selecting mock chat: %w", err)
			} else if errors.Is(err, errorhandling.ErrNoResultsFound) {
				err := chatcontrollers.InsertChat(dbutils.DB, userIds[currentUserId], chatData)
				if err != nil {
					return fmt.Errorf("failed to insert mock chat data: %w", err)
				}
				chat, err = chatcontrollers.SelectChat(dbutils.DB, userIds[currentUserId], followees.Followers[randomUser].FolloweeId)
				if err != nil {
					return fmt.Errorf("error selecting mock chat: %w", err)
				}
			}

			chatMessageData["chat_id"] = chat.ChatId
			chatMessageData["message"] = "Hey, how are ya?"

			for k := 0; k < rand.Intn(9)+1; k++ {
				err = chatcontrollers.InsertChatMessage(dbutils.DB, userIds[currentUserId], chatMessageData)
				if err != nil {
					return fmt.Errorf("failed to insert mock chat message data: %w", err)
				}
			}

			followees.Followers = append(followees.Followers[:randomUser], followees.Followers[randomUser+1:]...)

		}

	}
	return nil
}
