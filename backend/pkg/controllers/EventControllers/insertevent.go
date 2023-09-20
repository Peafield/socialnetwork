package eventcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"time"
)

func InsertEvent(db *sql.DB, userId string, eventData map[string]interface{}) error {
	args := make([]interface{}, 6)

	//create event id
	eventId, err := helpers.CreateUUID()
	if err != nil {
		return fmt.Errorf("failed to create event uuid: %w", err)
	}
	args[0] = eventId

	groupId, ok := eventData["group_id"].(string)
	if !ok {
		return fmt.Errorf("group id is not a string or doesn't exist")
	}
	args[1] = groupId

	creatorId, ok := eventData["creator_id"].(string)
	if !ok {
		return fmt.Errorf("creator id is not a string or doesn't exist")
	}
	args[2] = creatorId

	title, ok := eventData["title"].(string)
	if !ok {
		return fmt.Errorf("title is not a string or doesn't exist")
	}
	args[3] = title

	description, ok := eventData["description"].(string)
	if !ok {
		return fmt.Errorf("description is not a string or doesn't exist")
	}
	args[4] = description

	eventStartTime, ok := eventData["event_start_time"].(string)
	if !ok {
		return fmt.Errorf("event start time is not a string or doesn't exist")
	}
	formattedDOB, err := time.Parse("2006-01-02T15:04", eventStartTime)
	if err != nil {
		return fmt.Errorf("DOB string can't be parsed into time.Time: %w", err)
	}
	args[5] = formattedDOB

	err = crud.InteractWithDatabase(db, dbstatements.InsertGroupsEventsStmt, args)
	if err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}

	return nil
}
