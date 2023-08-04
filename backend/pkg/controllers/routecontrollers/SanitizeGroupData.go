package routecontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
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
	isCreator, err := IsGroupCreator(db, userId, groupId)
	if err != nil {
		return nil, fmt.Errorf("problem querying user/group relationship at SelectGroup: %v", err)
	}
	isMember, err := IsGroupMember(db, userId, groupId)
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
func UpdateGroup(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
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
func DeleteGroup(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
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
