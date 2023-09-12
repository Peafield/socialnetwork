package notificationcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
)

func InsertNewNotification(db *sql.DB, userId string, newNotificationData map[string]interface{}) error {
	args := make([]interface{}, 9)
	for i := range args {
		args[i] = ""
	}

	notificationId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create notification id: %w", err)
	}
	args[0] = notificationId

	args[1] = userId

	creatorId, ok := newNotificationData["receiver"].(string)
	if !ok {
		return fmt.Errorf("receiver is not a string or doesn't exist")
	}
	args[2] = creatorId

	groupId, ok := newNotificationData["group_id"].(string)
	if ok {
		args[3] = groupId
	}

	postId, ok := newNotificationData["post_id"].(string)
	if ok {
		args[4] = postId
	}

	eventId, ok := newNotificationData["event_id"].(string)
	if ok {
		args[5] = eventId
	}

	commentId, ok := newNotificationData["comment_id"].(string)
	if ok {
		args[6] = commentId
	}

	chatId, ok := newNotificationData["chat_id"].(string)
	if ok {
		args[7] = chatId
	}

	reactionType, ok := newNotificationData["action_type"].(string)
	if ok {
		args[8] = reactionType
	}

	err = crud.InteractWithDatabase(db, dbstatements.InsertNotificationsStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert notification into database: %w", err)
	}
	return nil
}
