package groupcontrollers

// const (
// 	attendeeInsertStmnt    = `INSERT INTO Groups_Members (event_id, attendee_id, attendee_status) VALUES (?, ?, ?)`
// 	oneAttendeeSelectStmnt = `SELECT * FROM Groups_Events_Attendees WHERE event_id = ? AND attendee_id = ?`
// 	allAttendeeSelectStmnt = `SELECT * FROM Groups_Events_Attendees WHERE event_id = ?`
// 	attendeeUpdateStmnt    = `UPDATE Groups_Events_Attendees SET attending_status = ? WHERE event_id = ? AND attendee_id = ? AND EXISTS (SELECT 1 FROM Groups_Events_attendees WHERE attendee_id = ?);`
// 	attendeeDeleteStmnt    = `DELETE FROM Groups_Events_Attendees WHERE event_id = ? AND attendee_id = ? AND EXISTS (SELECT 1 FROM Groups_Events_attendees WHERE attendee_id = ?);`
// )

// func InsertAttendee(db *sql.DB, attendeeId, eventId string, AffectedColumns map[string]interface{}) error {
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups", attendeeId, AffectedColumns["group_id"].(string))
// 	if !isMember {
// 		return errors.New("user is not part of the group. Therefore, not allowed to create an event")
// 	}

// 	Stmnt, err := db.Prepare(attendeeInsertStmnt)
// 	if err != nil {
// 		return errors.New("failed to create statement at insertevent")
// 	}

// 	Values := []interface{}{
// 		eventId,
// 		attendeeId,
// 		AffectedColumns["attendee_status"]}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return errors.New("failed to insert event into database")
// 	}

// 	return nil
// }

// func SelectAttendee(db *sql.DB, userId, eventId string, AffectedColumns map[string]interface{}) (interface{}, error) {
// 	//Check whether user is part of the desired group
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups", userId, AffectedColumns["group_id"].(string))
// 	if !isMember {
// 		return nil, errors.New("user is not part of the group. Therefore, not allowed to create an event")
// 	}

// 	var Statement string = allAttendeeSelectStmnt
// 	var Values []interface{} = []interface{}{eventId}

// 	attendeeId, ok := AffectedColumns["attendee_id"].(string)

// 	//If specific event ID is provided then select one 1
// 	if ok {
// 		Statement = oneEventSelectStmnt
// 		Values = append(Values, attendeeId)
// 	}
// 	//Prompt SELECT query
// 	result, err := crud.SelectFromDatabase(db, "Groups_Events", Statement, Values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func UpdateAttendee(db *sql.DB, userId, attendeeId, groupId, eventId string, attendingStatus bool) error {
// 	//check if user in question has the right to remove the attendee
// 	isCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	if !isCreator && userId != attendeeId {
// 		return errors.New("user has no rights to update attendee")
// 	}

// 	Values := []interface{}{attendingStatus, eventId, attendeeId, attendeeId}
// 	Stmnt, err := db.Prepare(attendeeUpdateStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func DeleteAttendee(db *sql.DB, userId, attendeeId, eventId, groupId string) error {
// 	//check if user in question has the right to remove the attendee
// 	isCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	if !isCreator && userId != attendeeId {
// 		return errors.New("user has no rights to remove attendee")
// 	}

// 	Values := []interface{}{eventId, attendeeId, attendeeId}
// 	Stmnt, err := db.Prepare(attendeeDeleteStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
