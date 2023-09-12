package notificationcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/models/dbmodels"
)

func UpdateAllNotificationsReadStatus(db *sql.DB, userId string, notifications *dbmodels.Notifications, readStatus int) error {
	var err error
	for _, n := range notifications.Notifications {
		err = crud.InteractWithDatabase(db, dbstatements.UpdateAllUserNotifications, []interface{}{readStatus, n.NotificationId})
		if err != nil {
			return fmt.Errorf("failed to update logged in status: %w", err)
		}
	}

	return nil
}
