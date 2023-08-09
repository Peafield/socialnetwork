package routecontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
)

const (
	insertMemberStmnt = `INSERT INTO Groups_Members (group_id, member_id) VALUES (?, ?)`
)

func InsertMember(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {

	// //check whether the map's keys match the expected parameters
	expectedParams := []string{"group_id"}
	exists := helpers.ValuesMapComparison(AffectedColumns, expectedParams)
	if !exists {
		return fmt.Errorf("parameters received = %v \n expected parameters = %v at InsertMember", AffectedColumns, expectedParams)
	}

	groupId := AffectedColumns["group_id"].(string)
	//check whether user has already received an invitation from the specified group
	//or whether the user is already a member

	//specify the required parameters to interact with the database
	Values := []interface{}{groupId, userId}
	Stmnt, err := db.Prepare(insertMemberStmnt)
	if err != nil {
		return err
	}

	//Prompt the Insert Query with its respective values
	crud.InteractWithDatabase(dbutils.DB, Stmnt, Values)

	return nil
}

func SelectMember() {}

func UpdateMember() {}

func DeleteMember() {}
