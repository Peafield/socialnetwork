package groupcontrollers

// const (
// 	eventInsertStmnt    = `INSERT INTO Groups_Events (event_id, group_id, creator_id, title, description, event_start_time) VALUES (?, ?, ?, ?, ?, ?)`
// 	allEventSelectStmnt = `SELECT * FROM Groups_Events WHERE group_id = ?`
// 	oneEventSelectStmnt = `SELECT * FROM Groups_Events WHERE group_id = ? AND event_id = ?`
// 	eventUpdateStmnt    = `UPDATE Groups_Events SET title = ?, description = ?, event_start_time = ? WHERE event_id = ?`
// 	eventDeleteStmnt    = `
// 	DELETE FROM Groups_Events WHERE event_id = ?;
// 	DELETE FROM Groups_Events_Attendees WHERE event_id = ?;
// 	`
// )

// func InsertEvent(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
// 	//check whether userId is a member of the group or the creator
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups", userId, AffectedColumns["group_id"].(string))
// 	if !isMember {
// 		return errors.New("user is not part of the group. Therefore, not allowed to create an event")
// 	}

// 	uuid, err := helpers.CreateUUID()
// 	if err != nil {
// 		return errors.New("failed to create uuid at insertevent")
// 	}

// 	Stmnt, err := db.Prepare(eventInsertStmnt)
// 	if err != nil {
// 		return errors.New("failed to create statement at insertevent")
// 	}

// 	Values := []interface{}{
// 		uuid,
// 		AffectedColumns["group_id"],
// 		userId,
// 		AffectedColumns["title"],
// 		AffectedColumns["description"],
// 		AffectedColumns["event_start_time"]}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return errors.New("failed to insert event into database")
// 	}

// 	return nil
// }

// func SelectEvent(db *sql.DB, userId, groupId string, AffectedColumns map[string]interface{}) (interface{}, error) {
// 	//Check whether user is part of the desired group
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups", userId, AffectedColumns["group_id"].(string))
// 	if !isMember {
// 		return nil, errors.New("user is not part of the group. Therefore, not allowed to create an event")
// 	}

// 	var Statement string = allEventSelectStmnt
// 	var Values []interface{} = []interface{}{groupId}

// 	eventId, ok := AffectedColumns["event_id"]

// 	//If specific event ID is provided then select one 1
// 	if ok {
// 		Statement = oneEventSelectStmnt
// 		Values = append(Values, eventId)
// 	}

// 	//Prompt SELECT query
// 	result, err := crud.SelectFromDatabase(db, "Groups_Events", Statement, Values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func UpdateEvent(db *sql.DB, userId, groupId, eventId string, AffectedColumns map[string]interface{}) error {
// 	//check whether user has permission to update event
// 	//should be either event creator or group creator.
// 	isGroupCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	isEventCreator := dbutils.IsEventCreator(db, userId, eventId)

// 	if !isGroupCreator && !isEventCreator {
// 		return errors.New("user has no permission to update event")
// 	}

// 	//title, description, event_start_time
// 	Values := []interface{}{AffectedColumns["title"], AffectedColumns["description"], AffectedColumns["event_start_time"], eventId}
// 	Stmnt, err := db.Prepare(eventUpdateStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func DeleteEvent(db *sql.DB, userId, groupId, eventId string) error {
// 	//check whether user has permission to update event
// 	//should be either event creator or group creator.
// 	isGroupCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	isEventCreator := dbutils.IsEventCreator(db, userId, eventId)

// 	if !isGroupCreator && !isEventCreator {
// 		return errors.New("user has no permission to delete event")
// 	}

// 	//Also delete everything related to the group
// 	Stmnt, err := db.Prepare(eventDeleteStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, []interface{}{eventId})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
