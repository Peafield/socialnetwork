package routecontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
)

func SignOutUser(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
	conditions := make(map[string]interface{})
	conditions["user_id"] = userId

	//make sure immutable parameters are not trying to be changed
	MutableParameters := []string{"is_logged_in"}

	//check whether the mutable parameters are deeply equal to the map
	dataContainsImmutableParameter := !helpers.ValuesMapComparison(AffectedColumns, MutableParameters)

	if dataContainsImmutableParameter {
		return fmt.Errorf("error trying to update user immutable parameter at SignOutUser")
	}

	AffectedColumns["is_logged_in"] = 0
	err := crud.UpdateDatabaseRow(dbutils.DB, "Users", conditions, AffectedColumns)

	if err != nil {
		return err
	}

	return nil
}
