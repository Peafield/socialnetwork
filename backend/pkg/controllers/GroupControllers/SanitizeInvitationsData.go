package groupcontrollers

// const (
// 	invitationInsertStmnt    = `INSERT INTO Groups_Invitations (group_id, user_id, is_invited) VALUES (?, ?, ?)`
// 	oneInvitationSelectStmnt = `SELECT * FROM Groups_Invitations WHERE user_id = ? AND group_id = ?`
// 	allInvitationSelectStmnt = `SELECT * FROM Groups_Invitations WHERE user_id = ?`
// 	//how can you update an invitation?
// 	invitationUpdateStmnt = `UPDATE Groups_Invitations SET user_id = ? WHERE user_id = ? AND group_id = ?`
// 	invitationDeleteStmnt = `DELETE FROM Groups_Invitations WHERE group_id = ? AND user_id = ?`
// )

// func InsertInvitation(db *sql.DB, groupId, userId string, isInvited bool) error {
// 	//before inserting we need to check wether the user is part of the group already or not.
// 	//OR if the user already received an invitation
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups", userId, groupId)
// 	isInvitationReceived := dbutils.DoesGroupRowExist(db, "Groups_Invitations", userId, groupId)

// 	switch {
// 	case isMember:
// 		return errors.New("user is already a member of the group")

// 	case isInvitationReceived:
// 		return errors.New("user is already received an invitation from the group")

// 	default:
// 		Values := []interface{}{groupId, userId, isInvited}
// 		Stmnt, err := db.Prepare(invitationInsertStmnt)
// 		if err != nil {
// 			return err
// 		}
// 		err = crud.InteractWithDatabase(db, Stmnt, Values)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func SelectInvitation(db *sql.DB, userId string, AffectedColumns map[string]interface{}) (interface{}, error) {

// 	var Statement string = allAttendeeSelectStmnt
// 	var Values []interface{} = []interface{}{userId}

// 	groupId, ok := AffectedColumns["group_id"].(string)

// 	//If specific group ID is provided then select one 1
// 	if ok {
// 		Statement = oneEventSelectStmnt
// 		Values = append(Values, groupId)
// 	}
// 	//Prompt SELECT query
// 	result, err := crud.SelectFromDatabase(db, "Groups_Invitations", Statement, Values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil

// }

// // no clue how to update an invitation
// func UpdateInvitation() {

// }

// func DeleteInvitation(db *sql.DB, userId, groupId string) error {

// 	//Prepare values
// 	Values := []interface{}{groupId, userId}
// 	Stmnt, err := db.Prepare(invitationDeleteStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	//Prompt DELETE query
// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
