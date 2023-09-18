package groupcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectAllGroupMembers(db *sql.DB, groupId string) (*dbmodels.GroupMembers, error) {
	groupMembersData, err := crud.SelectFromDatabase(db, "Groups_Members", dbstatements.SelectAllGroupMembersStmt, []interface{}{groupId})
	if err != nil {
		return nil, fmt.Errorf("could not select group members: %w", err)
	}

	groupMembers := &dbmodels.GroupMembers{}

	for _, v := range groupMembersData {
		if groupMember, ok := v.(*dbmodels.GroupMember); ok {
			groupMembers.GroupMembers = append(groupMembers.GroupMembers, *groupMember)
		} else {
			return nil, fmt.Errorf("could not assert group member type")
		}
	}

	return groupMembers, nil
}
