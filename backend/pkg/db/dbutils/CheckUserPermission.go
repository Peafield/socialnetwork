package dbutils

import (
	"database/sql"
	"fmt"
)

func DoesGroupRowExist(db *sql.DB, tableName, userId, groupId string) bool {

	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM ? WHERE user_id = ? AND group_id = ?)"
	//not sure whether 0 should be an int or a string
	err := db.QueryRow(query, tableName, userId, groupId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func IsGroupCreator(db *sql.DB, userId, groupId string) bool {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups WHERE creator_id = ? AND group_id = ?)"

	err := db.QueryRow(query, userId, groupId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
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

func IsEventCreator(db *sql.DB, userId, eventId string) bool {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups_Events WHERE creator_id = ? AND event_id = ?)"

	err := db.QueryRow(query, userId, eventId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func IsEventAttendee(db *sql.DB, userId, eventId string) bool {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM Groups_Events_Attendees WHERE attendee_id = ? AND event_id = ?)"

	err := db.QueryRow(query, userId, eventId).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
