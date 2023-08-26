package groupcontrollers

// const (
// 	groupInsertStmnt = `INSERT INTO Groups (group_id, title, description, creator_id) VALUES (?, ?, ?, ?)`
// 	groupUpdateStmnt = `UPDATE Groups SET title = ?, description = ? WHERE group_id = ?`
// 	groupSelectStmnt = `SELECT * FROM Groups WHERE group_id = ?`

// 	groupDeleteStatements = `
// DELETE FROM Groups WHERE group_id = ?;
// DELETE FROM Groups_Members WHERE group_id = ?;
// DELETE FROM Groups_Events WHERE group_id = ?;
// DELETE FROM Groups_Events_Attendees WHERE group_id = ?;
// DELETE FROM Groups_Invitations WHERE group_id = ?;
// `
// 	postDeleteStatements = `
// DELETE FROM Posts WHERE post_id = ?;
// DELETE FROM Comments WHERE post_id = ?;
// DELETE FROM Reactions WHERE post_id = ?;
// `
// )

// /**/
// func InsertGroup(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {

// 	//make sure immutable parameters are not trying to be changed
// 	expectedParams := []string{"title", "description"}

// 	//check whether the map's keys match the expected parameters
// 	exists := helpers.MapKeyContains(AffectedColumns, expectedParams)

// 	if !exists {
// 		return fmt.Errorf("parameters received = %v expected parameters = %v at InsertGroup", AffectedColumns, expectedParams)
// 	}

// 	uuid, err := helpers.CreateUUID()
// 	if err != nil {
// 		return nil
// 	}

// 	//prepare the arguments for InteractWithDatabase
// 	Values := []interface{}{uuid, AffectedColumns["title"], AffectedColumns["description"], userId}
// 	Stmnt, err := db.Prepare(groupInsertStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// /**/
// func SelectGroup(db *sql.DB, userId string, Conditions map[string]interface{}) ([]interface{}, error) {

// 	groupId, exists := Conditions["group_id"].(string)
// 	if !exists {
// 		return nil, fmt.Errorf("invalid parameters received at SelectGroup")
// 	}

// 	// Check User/Group relationship
// 	// If user isn't allowed access to group then return an error
// 	isCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	isMember := dbutils.DoesGroupRowExist(db, "Groups_Members", userId, groupId)

// 	if !isMember && !isCreator {
// 		return nil, errors.New("user has no rights to access group in question")
// 	}

// 	//Prompt Select Query
// 	queryResult, err := crud.SelectFromDatabase(db, "Groups", groupSelectStmnt, []interface{}{groupId})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return queryResult, nil
// }

// /**/
// func UpdateGroup(db *sql.DB, userId, groupId string, AffectedColumns map[string]interface{}) error {

// 	//check permission level of user
// 	isPermitted, err := dbutils.IsPermitted(dbutils.DB, userId, groupId, "Groups")
// 	if !isPermitted || err != nil {
// 		return fmt.Errorf("user isn't permitted to update Group: %v", groupId)

// 	}

// 	//check whether the map meets the expected parameters
// 	expectedParameters := []string{"title", "description"}
// 	found := helpers.FoundParameters(AffectedColumns, expectedParameters)
// 	if !found {
// 		return fmt.Errorf("expected parameters = %v \nreceived = %v  at UpdateGroup", expectedParameters, AffectedColumns)
// 	}

// 	//prepare the arguments for InteractWithDatabase
// 	Values := []interface{}{AffectedColumns["title"], AffectedColumns["description"], groupId}
// 	Stmnt, err := db.Prepare(groupUpdateStmnt)
// 	if err != nil {
// 		return err
// 	}

// 	//Execute the update statement
// 	err = crud.InteractWithDatabase(db, Stmnt, Values)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// /**/
// func DeleteGroup(db *sql.DB, userId, groupId string) error {

// 	isCreator := dbutils.IsGroupCreator(db, userId, groupId)
// 	if !isCreator {
// 		return errors.New("user doesn't have permission to delete group")
// 	}

// 	//delete group and all related info: events, members, posts, comments, likes&dislikes
// 	//by initializing a transaction to handle all operations at once. If one fails then the transaction will never take effect

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return errors.New("failed to begin transaction at DeleteGroup: " + err.Error())
// 	}

// 	//supply the group Id 4 times as required by the groupDeleteStatement
// 	_, err = tx.Exec(groupDeleteStatements, groupId, groupId, groupId, groupId)
// 	if err != nil {
// 		tx.Rollback()
// 		return errors.New("failed to execute transaction at DeleteGroup: " + err.Error())
// 	}

// 	//select all posts related to the group in question
// 	PostsData, err := crud.SelectFromDatabase(dbutils.DB, "Posts", "SELECT * FROM Posts WHERE group_id = ?", []interface{}{groupId})
// 	if err != nil {
// 		return err
// 	}

// 	//for each post delete its related comments and reactions
// 	for _, post := range PostsData {
// 		postId := (post.(dbmodels.Post)).PostId

// 		//supply the group ID once and the post ID twice as required by the postDeleteStatement
// 		_, err = tx.Exec(postDeleteStatements, postId, postId, postId)
// 		if err != nil {
// 			tx.Rollback()
// 			return errors.New("failed to execute transaction at DeleteGroup: " + err.Error())
// 		}

// 	}

// 	//finally, commit every query requested since the outset of the transaction
// 	err = tx.Commit()
// 	if err != nil {
// 		return errors.New("failed to commit transaction at DeleteGroup: " + err.Error())

// 	}

// 	return nil
// }
