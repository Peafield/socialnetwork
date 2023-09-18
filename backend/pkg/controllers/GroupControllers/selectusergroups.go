package groupcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func GetUserGroups(db *sql.DB, userId string) (*dbmodels.Groups, error) {
	groupsData, err := crud.SelectFromDatabase(db, "Groups", dbstatements.SelectUserGroupsStmt, []interface{}{userId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("error selecting groups %w", err)
	}

	groups := &dbmodels.Groups{}
	for _, v := range groupsData {
		if group, ok := v.(*dbmodels.Group); ok {
			groups.Groups = append(groups.Groups, *group)
		} else {
			return nil, fmt.Errorf("failed to assert group data")
		}
	}

	return groups, nil
}
