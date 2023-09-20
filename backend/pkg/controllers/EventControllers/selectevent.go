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

func SelectAllGroupEvents(db *sql.DB, userId string, groupId string) (*dbmodels.GroupEvents, error) {
	eventsData, err := crud.SelectFromDatabase(db, "Groups_Events", dbstatements.SelectAllGroupEventsStmt, []interface{}{groupId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("failed to select events data: %w", err)
	}

	events := &dbmodels.GroupEvents{}
	for _, v := range eventsData {
		if event, ok := v.(*dbmodels.GroupEvent); ok {
			events.GroupEvents = append(events.GroupEvents, *event)
		} else {
			return nil, fmt.Errorf("failed to assert group event data")
		}
	}

	return events, nil
}
