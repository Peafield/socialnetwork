package usercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func UpdateLoggedInStatus(db *sql.DB, userId string, loggedInStatus int) error {
	var statement *sql.Stmt

	switch loggedInStatus {
	case 0:
		statement = dbstatements.UpdateUserLoggedOut
		break
	case 1:
		statement = dbstatements.UpdateUserLoggedIn
		break
	default:
		return fmt.Errorf("logged in status is not a valid value")
	}

	err := crud.InteractWithDatabase(db, statement, []interface{}{userId})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update logged in status: %w", err)
	}
	return nil
}
