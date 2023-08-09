package routecontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

/**/
func InsertGroup(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
	// conditions := make(map[string]interface{})
	// conditions["user_id"] = userId

	//make sure immutable parameters are not trying to be changed
	expectedParams := []string{"title", "description"}

	//check whether the map's keys match the expected parameters
	exists := helpers.MapKeyContains(AffectedColumns, expectedParams)

	if !exists {
		return fmt.Errorf("parameters received = %v expected parameters = %v at InsertGroup", AffectedColumns, expectedParams)
	}

	uuid, err := helpers.CreateUUID()
	if err != nil {
		return nil
	}
	insertValues := []interface{}{uuid, AffectedColumns["title"], AffectedColumns["description"], userId}
	crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertGroupsStmt, insertValues)

	return nil
}

/**/
func SelectGroup(db *sql.DB, userId string, Conditions map[string]interface{}) ([]interface{}, error) {

	groupId, exists := Conditions["group_id"].(string)
	if !exists {
		return nil, fmt.Errorf("invalid parameters received at SelectGroup")
	}

	// Check User/Group relationship
	// If user isn't allowed access to group then return an error
	isCreator, err := dbutils.IsGroupCreator(db, userId, groupId)
	if err != nil {
		return nil, fmt.Errorf("problem querying user/group relationship at SelectGroup: %v", err)
	}
	isMember, err := dbutils.IsGroupMember(db, userId, groupId)
	if err != nil {
		return nil, fmt.Errorf("problem querying user/group relationship at SelectGroup: %v", err)
	}

	//if client isn't related to the group in question
	if !isMember && !isCreator {
		return nil, errors.New("user has no rights to access group in question")
	}

	//Prompt Select Query
	query := `SELECT * FROM Groups WHERE group_id = ?`

	queryResult, err := crud.SelectFromDatabase(db, "Groups", query, []interface{}{groupId})

	if err != nil {
		return nil, errors.New("user has no rights to access group in question")
	}
	return queryResult, nil
}

/**/
func UpdateGroup(db *sql.DB, userId, groupId string, AffectedColumns map[string]interface{}) error {

	//check permission level of user
	isPermitted, err := dbutils.IsPermitted(dbutils.DB, userId, groupId, "Groups")
	if !isPermitted || err != nil {
		return fmt.Errorf("user isn't permitted to update Group: %v", groupId)

	}

	//should maybe ignore immutable parameters received instead of returning errors
	immutableParameters := []string{"group_id", "creator_id", "creation_date"}

	//check whether the map's keys match the expected parameters
	exists := helpers.MapKeyContains(AffectedColumns, immutableParameters)

	if exists {
		return fmt.Errorf("parameters received = %v expected parameters = %v at InsertGroup", AffectedColumns, immutableParameters)
	}

	// updateErr := []interface{}{uuid, AffectedColumns["title"], AffectedColumns["description"], userId}
	err = crud.UpdateDatabaseRow(dbutils.DB, "Groups", map[string]interface{}{"group_id": groupId}, AffectedColumns)
	if err != nil {
		return err
	}
	return nil
}

/**/
func DeleteGroup(db *sql.DB, userId, groupId string) error {

	isCreator, err := dbutils.IsGroupCreator(db, userId, groupId)
	if !isCreator || err != nil {
		return fmt.Errorf("user doesn't have permission to delete group: %s", groupId)
	}

	//delete group and all related info: events, members, posts, comments, likes&dislikes
	err = crud.DeleteFromDatabase(dbutils.DB, "Groups", map[string]interface{}{"group_id": groupId})
	if err != nil {
		return err
	}
	err = crud.DeleteFromDatabase(dbutils.DB, "Groups_Members", map[string]interface{}{"group_id": groupId})
	if err != nil {
		return err
	}
	err = crud.DeleteFromDatabase(dbutils.DB, "Groups_Events", map[string]interface{}{"group_id": groupId})
	if err != nil {
		return err
	}
	err = crud.DeleteFromDatabase(dbutils.DB, "Groups_Events_Attendees", map[string]interface{}{"group_id": groupId})
	if err != nil {
		return err
	}

	//select the post id and comments id to delete the likes and dislikes associated with it
	PostsData, err := crud.SelectFromDatabase(dbutils.DB, "Posts", "SELECT * FROM Posts WHERE group_id = ?", []interface{}{groupId})
	if err != nil {
		return err
	}
	//for each post delete the related comments and likes/dislikes associated with it
	for _, post := range PostsData {
		//assign a variable for the postId
		postId := (post.(dbmodels.Post)).PostId
		err = crud.DeleteFromDatabase(dbutils.DB, "Posts", map[string]interface{}{"post_id": postId})
		if err != nil {
			return err
		}
		err = crud.DeleteFromDatabase(dbutils.DB, "Comments", map[string]interface{}{"post_id": postId})
		if err != nil {
			return err
		}
		err = crud.DeleteFromDatabase(dbutils.DB, "reactions", map[string]interface{}{"post_id": postId})
		if err != nil {
			return err
		}

	}

	err = crud.DeleteFromDatabase(dbutils.DB, "Posts", map[string]interface{}{"group_id": groupId})
	if err != nil {
		return err
	}

	return nil
}
