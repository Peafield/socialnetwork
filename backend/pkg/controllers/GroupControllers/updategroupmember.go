package groupcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
)

func UpdateGroupMember(db *sql.DB, userId string, updateGroupMemberData map[string]interface{}) error {
	args := make([]interface{}, 3)

	memberId, ok := updateGroupMemberData["member_id"].(string)
	if !ok {
		return fmt.Errorf("member id is not a string or doesn't exist")
	}

	groupId, ok := updateGroupMemberData["group_id"].(string)
	if !ok {
		return fmt.Errorf("group id is not a string or doesn't exist")
	}

	accepted, ok := updateGroupMemberData["accepted"].(bool)
	if !ok {
		return fmt.Errorf("accepted boolean is not a bool or doesn't exist")
	} else {
		if accepted {
			args[0] = 1
		} else {
			args[0] = 0
			// delete group member?
		}
	}

	args[1] = memberId

	args[2] = groupId

	err := crud.InteractWithDatabase(db, dbstatements.UpdateGroupMemberStatusStmt, args)
	if err != nil {
		return fmt.Errorf("could not update group member status: %w", err)
	}

	return nil

}
