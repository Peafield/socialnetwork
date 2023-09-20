package eventcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func InsertEventAttendee(db *sql.DB, userId string, eventAttendeeData map[string]interface{}) error {
	args := make([]interface{}, 3)

	eventId, ok := eventAttendeeData["event_id"].(string)
	if !ok {
		return fmt.Errorf("event id is not a string or doesn't exist")
	}
	args[0] = eventId

	attendee_id, ok := eventAttendeeData["attendee_id"].(string)
	if !ok {
		return fmt.Errorf("attendee id is not a string or doesn't exist")
	}
	args[1] = attendee_id

	attendingStatus, ok := eventAttendeeData["attending_status"].(float64)
	if !ok {
		return fmt.Errorf("event id is not a float64 or doesn't exist")
	}
	args[2] = int(attendingStatus)

	err := crud.InteractWithDatabase(db, dbstatements.InsertGroupsEventsAttendees, args)
	if err != nil {
		return fmt.Errorf("failed to insert group event attendee: %w", err)
	}

	return nil

}
