package notificationcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func DeleteNotification(db *sql.DB, userId string, notificationData map[string]interface{}) error {
	notificationId, ok := notificationData["notification_id"].(string)
	if !ok {
		return fmt.Errorf("notification id is not a string or doesnt exist")
	}

	err := crud.InteractWithDatabase(db, dbstatements.DeleteNotificationStmt, []interface{}{notificationId})
	if err != nil {
		return fmt.Errorf("could not delete notfication: %w", err)
	}

	return nil
}
