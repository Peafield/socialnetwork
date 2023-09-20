package eventcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func UpdateAttendeeStatus(db *sql.DB, userId string, updateAttendeeData map[string]interface{}) error {
	args := make([]interface{}, 3)

	eventId, ok := updateAttendeeData["event_id"].(string)
	if !ok {
		return fmt.Errorf("event id is not a string or doesn't exist")
	}

	attendee_id, ok := updateAttendeeData["attendee_id"].(string)
	if !ok {
		return fmt.Errorf("attendee id is not a string or doesn't exist")
	}

	attendingStatus, ok := updateAttendeeData["attending_status"].(float64)
	if !ok {
		return fmt.Errorf("event id is not a float64 or doesn't exist")
	}

	args[0] = int(attendingStatus)
	args[1] = attendee_id
	args[2] = eventId

	err := crud.InteractWithDatabase(db, dbstatements.UpdateAttendeeStatus, args)
	if err != nil {
		return fmt.Errorf("failed to insert group event attendee")
	}

	return nil
}
