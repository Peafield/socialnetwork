package eventcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectAllEventAttendees(db *sql.DB, userId string, eventId string) (*dbmodels.GroupEventAttendees, error) {
	attendeesData, err := crud.SelectFromDatabase(db, "Groups_Events_Attendees", dbstatements.SelectAllEventAttendeesStmt, []interface{}{eventId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select attendees data: %w", err)
	}

	attendees := &dbmodels.GroupEventAttendees{}
	for _, v := range attendeesData {
		if eventAttendee, ok := v.(*dbmodels.GroupEventAttendee); ok {
			attendees.GroupEventAttendees = append(attendees.GroupEventAttendees, *eventAttendee)
		} else {
			return nil, fmt.Errorf("failed to assert group event data")
		}
	}

	return attendees, nil
}
