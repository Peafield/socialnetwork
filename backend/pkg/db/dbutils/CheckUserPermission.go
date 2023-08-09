package dbutils

import (
	"database/sql"
	"fmt"
)

func IsGroupMember(db *sql.DB, userId, groupId string) (bool, error) {

	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups_Members WHERE user_id = ?)"
	//not sure whether 0 should be an int or a string
	err := db.QueryRow(query, userId).Scan(&exists)
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

func IsPermitted(db *sql.DB, userId, groupId, tableName string) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = ? AND group_id = ? AND NOT permission_level = 0)", tableName)

	err := db.QueryRow(query, userId, groupId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
