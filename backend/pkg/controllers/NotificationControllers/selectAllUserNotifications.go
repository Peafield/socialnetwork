package notificationcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectAllUserNotifications(db *sql.DB, userId string) (*dbmodels.Notifications, error) {
	notificationsData, err := crud.SelectFromDatabase(db, "Notifications", dbstatements.SelectAllUserNotifications, []interface{}{userId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select all user notifications from database: %w", err)
	}

	notifications := &dbmodels.Notifications{}
	for _, v := range notificationsData {
		if notification, ok := v.(*dbmodels.Notification); ok {
			notifications.Notifications = append(notifications.Notifications, *notification)
		} else {
			return nil, fmt.Errorf("failed to assert post data")
		}
	}

	return notifications, nil
}
