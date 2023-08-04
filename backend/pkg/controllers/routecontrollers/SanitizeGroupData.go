package routecontrollers

import (
	"database/sql"
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
func SelectGroup(db *sql.DB, userId string, Conditions map[string]interface{}) error {

	//make sure immutable parameters are not trying to be changed
	expectedParams := []string{"title", "description"}

	//check whether the map's keys match the expected parameters
	exists := helpers.MapKeyContains(Conditions, expectedParams)

	if !exists {
		return fmt.Errorf("parameters received = %v expected parameters = %v at InsertGroup", Conditions, expectedParams)
	}

	// uuid, err := helpers.CreateUUID()
	// if err != nil {
	// 	return nil
	// }
	// insertValues := []interface{}{uuid, C["title"], AffectedColumns["description"], userId}
	// crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertGroupsStmt, insertValues)

	return nil
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
