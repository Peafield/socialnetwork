package groupcontrollers

// const (
// 	memberInsertStmnt = `
// 	DELETE FROM Groups_Invitations WHERE group_id = ? AND user_id = ?;
// 	INSERT INTO Groups_Members (group_id, member_id) VALUES (?, ?);
// 	`
// 	oneMemberSelectStmnt = `SELECT * FROM Groups_Members WHERE group_id = ? AND user_id = ?`
// 	allMemberSelectStmnt = `SELECT * FROM Groups_Members WHERE group_id = ?`
// 	memberUpdateStmnt    = `UPDATE Groups_Members SET permission_level = ? WHERE group_id = ? AND user_id = ?`
// 	memberDeleteStmnt    = `DELETE FROM Groups_Members WHERE group_id = ? AND user_id = ?`
// )

// func InsertMember(db *sql.DB, userId string, groupId string) error {

// 	//check whether user has already received an invitation from the specified group#
// 	//or whether the user is already a member
// 	isInvited := dbutils.DoesGroupRowExist(db, "Groups_Invitations", userId, groupId)
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups_Members", userId, groupId)

// 	if !isInvited && !isMember {
// 		return fmt.Errorf("no invitation found for the user in relation to the requested group")
// 	}

// 	if isInvited && isMember {
// 		//delete invitation from database
// 		// err := DeleteInvitation()
// 		return fmt.Errorf("user is already a member to the specified group")
// 	}

// 	//specify the required parameters to interact with the database
// 	//(DELETE invite and INSERT user into database)
// 	Values := []interface{}{groupId, userId, groupId, userId}
// 	Stmnt, err := db.Prepare(memberInsertStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	//Prompt the Insert Query with its respective values
// 	err = crud.InteractWithDatabase(dbutils.DB, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SelectMember(db *sql.DB, groupId string, Conditions map[string]interface{}) (interface{}, error) {

// 	var Statement string = allMemberSelectStmnt
// 	var Values []interface{} = []interface{}{groupId}

// 	//select 1 or multiple members?
// 	//if no parameters are provided then select all members
// 	userId, exists := Conditions["user_id"]
// 	if !exists {
// 		Statement = oneMemberSelectStmnt
// 		Values = append(Values, userId)
// 	}

// 	result, err := crud.SelectFromDatabase(db, "Groups_Members", Statement, Values)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// // this one might be useless
// func UpdateMember(db *sql.DB, adminId string, AffectedColumns map[string]interface{}) error {
// 	//UpdateMember involves modifying the state of a member in the group by an admin
// 	//might focus on transfering ownership to other members by the owner himself
// 	//might involve the promotion or demotion of a member by the owner himself

// 	//select the permission level of the user relative to the group
// 	//if the permission	is == 0 => return an error

// 	//if the permission level is > 0
// 	isCreator := dbutils.IsGroupCreator(db, adminId, AffectedColumns["group_id"].(string))
// 	if !isCreator {
// 		return errors.New("user doesn't have permissions to update members")
// 	}

// 	Values := []interface{}{AffectedColumns["permission_level"], AffectedColumns["group_id"], AffectedColumns["user_id"]}
// 	Stmnt, err := db.Prepare(memberUpdateStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}
// 	//must be able to transfer ownership to someone else

// 	return nil
// }
// func DeleteMember(db *sql.DB, adminId, memberId, groupId string) error {
// 	//a delete member involves deleting a member from the group
// 	//either from the user himself or from a non-regular member
// 	isCreator := dbutils.IsGroupCreator(db, adminId, groupId)
// 	if !isCreator {
// 		return errors.New("user doesn't have permissions to delete members")
// 	}

// 	Values := []interface{}{groupId, memberId}
// 	Stmnt, err := db.Prepare(memberDeleteStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
