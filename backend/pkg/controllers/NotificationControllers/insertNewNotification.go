package notificationcontrollers

import (
	"database/sql"
	"fmt"
	"log"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertNewNotification(db *sql.DB, userId string, newNotificationData map[string]interface{}) error {
	log.Println(newNotificationData)
	args := make([]interface{}, 9)

	notificationId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create notification id: %w", err)
	}
	args[0] = notificationId

	args[1] = userId

	creatorId, ok := newNotificationData["creatorId"].(string)
	if ok {
		args[2] = creatorId
	}

	notificationTypeId, ok := newNotificationData["notificationTypeId"].(string)
	if !ok {
		return fmt.Errorf("notificationTypeId is not a string")
	}

	reactionOn, ok := newNotificationData["reactionOn"].(string)
	if ok {
		switch reactionOn {
		case "group":
			args[3] = notificationTypeId
		case "post":
			args[4] = notificationTypeId
		case "event":
			args[5] = notificationTypeId
		case "comment":
			args[6] = notificationTypeId
		case "chat":
			args[7] = notificationTypeId
		}
	}

	reactionType, ok := newNotificationData["reactionType"].(string)
	if ok {
		args[8] = reactionType
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertNotificationsStmt, args)
	if err != nil {
		log.Printf("IWD error: %v", err)
		return fmt.Errorf("failed to insert notification into database: %w", err)
	}
	return nil
}
