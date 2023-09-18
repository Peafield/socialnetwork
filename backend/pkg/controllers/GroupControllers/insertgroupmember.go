package groupcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func InsertGroupCreatorMember(db *sql.DB, userId string, groupId string) error {
	args := make([]interface{}, 4)

	args[0] = groupId

	args[1] = userId

	//set request pending
	args[2] = 0

	//set permission level
	args[3] = 2

	err := crud.InteractWithDatabase(db, dbstatements.InsertGroupsMembersStmt, args)
	if err != nil {
		return fmt.Errorf("could not insert group member: %w", err)
	}

	return nil
}

func InsertGroupMembersAsInvitees(db *sql.DB, groupMemberData map[string]interface{}) error {
	args := make([]interface{}, 4)

	groupId, ok := groupMemberData["group_id"].(string)
	if !ok {
		return fmt.Errorf("group id is not a string or doesn't exist")
	}
	args[0] = groupId

	invitees, ok := groupMemberData["invitees_ids"].([]interface{})
	if !ok {
		return fmt.Errorf("invitees ids is not an array of string or doesn't exist")
	}

	//set request pending
	args[2] = 1

	//set permission level
	args[3] = 0

	for _, v := range invitees {
		args[1] = v

		err := crud.InteractWithDatabase(db, dbstatements.InsertGroupsMembersStmt, args)
		if err != nil {
			return fmt.Errorf("could not insert group member: %w", err)
		}
	}

	return nil
}

func InsertGroupMemberHasRequested(db *sql.DB, userId string, groupMemberData map[string]interface{}) error {
	args := make([]interface{}, 4)

	groupId, ok := groupMemberData["group_id"].(string)
	if !ok {
		return fmt.Errorf("group id is not a string or doesn't exist")
	}
	args[0] = groupId

	args[1] = userId

	args[2] = 2

	args[3] = 0

	err := crud.InteractWithDatabase(db, dbstatements.InsertGroupsMembersStmt, args)
	if err != nil {
		return fmt.Errorf("could not insert group member: %w", err)
	}

	return nil
}
