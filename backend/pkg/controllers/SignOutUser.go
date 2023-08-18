package controllers

import (
	"database/sql"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
)

func SignOutUser(db *sql.DB, userId string, AffectedColumns map[string]interface{}) error {
	args := []interface{}{userId}

	err := crud.InteractWithDatabase(dbutils.DB, dbstatements.UpdateUserLoggedOut, args)

	if err != nil {
		return err
	}

	return nil
}
