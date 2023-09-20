package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
)

func CreateFakeFollowers(db *sql.DB) error {
	followerAmount := 10
	followData := map[string]interface{}{}
	updateFollowdata := map[string]interface{}{}

	userIds, err := GetAllUserIDs(db)
	if err != nil {
		return fmt.Errorf("failed to get all user id's when faking follows: %w", err)
	}

	numOfUsers := len(userIds)
	usedUsers := []int{}

	for i := 0; i < numOfUsers; i++ {
		currentUserId := i
		usedUsers = append(usedUsers, currentUserId)
		for j := 0; j < followerAmount; j++ {

			randomUser := getRandomIntWithBlacklist(0, numOfUsers-1, usedUsers)
			usedUsers = append(usedUsers, randomUser)
			followData["followee_id"] = userIds[randomUser]
			updateFollowdata["follower_id"] = userIds[currentUserId]
			updateFollowdata["following_status"] = float64(1)

			err := followercontrollers.FollowUser(db, userIds[currentUserId], followData)
			if err != nil {
				return fmt.Errorf("failed to insert mock follow data: %w", err)
			}

			err = followercontrollers.UpdateFollowingStatus(db, userIds[randomUser], updateFollowdata)
			if err != nil {
				return fmt.Errorf("failed to update mock follow data: %w", err)
			}
		}
		usedUsers = []int{}
	}
	return nil
}

func getRandomIntWithBlacklist(min int, max int, blacklisted []int) int {

	// if blacklisted is/can be large, you might want to think about caching it
	excluded := map[int]bool{}
	for _, x := range blacklisted {
		excluded[x] = true
	}

	// loop until an n is generated that is not in the blacklist
	for {
		n := min + rand.Intn(max+1) // yields n such that min <= n <= max
		if !excluded[n] {
			return n
		}
	}

}
