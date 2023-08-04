package routecontrollers

import (
	"database/sql"
)

func IsGroupMember(db *sql.DB, userId, groupId string) (bool, error) {

	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups_Members WHERE user_id = ? AND request_pending = ?)"
	//not sure whether 0 should be an int or a string
	err := db.QueryRow(query, userId, `0`).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func IsGroupCreator(db *sql.DB, userId, groupId string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups WHERE creator_id = ? AND group_id = ?)"

	err := db.QueryRow(query, userId, groupId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
