package routecontrollers

import (
	"database/sql"
)

func InsertMember(db *sql.DB, AffectedColumns map[string]interface{}) error {

	// //make sure immutable parameters are not trying to be changed
	// expectedParams := []string{"user_id", "group_id"}

	// //check whether the map's keys match the expected parameters
	// exists := helpers.MapKeyContains(AffectedColumns, expectedParams)

	// if !exists {
	// 	return fmt.Errorf("parameters received = %v expected parameters = %v at InsertGroup", AffectedColumns, expectedParams)
	// }

	// uuid, err := helpers.CreateUUID()
	// if err != nil {
	// 	return nil
	// }
	// insertValues := []interface{}{uuid, AffectedColumns["title"], AffectedColumns["description"], userId}
	// crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertGroupsStmt, insertValues)

	return nil
}
