package groupcontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/models/dbmodels"
)

func GetGroupByTitle(db *sql.DB, title string) (*dbmodels.Group, error) {
	groupData, err := crud.SelectFromDatabase(db, "Groups", dbstatements.SelectGroupByTitleStmt, []interface{}{title})
	if err != nil {
		return nil, fmt.Errorf("error selecting group: %w", err)
	}

	group, ok := groupData[0].(*dbmodels.Group)
	if !ok {
		return nil, fmt.Errorf("could not assert group type")
	}

	return group, nil
}

func GetGroupByID(db *sql.DB, groupId string) (*dbmodels.Group, error) {
	groupData, err := crud.SelectFromDatabase(db, "Groups", dbstatements.SelectGroupByIDStmt, []interface{}{groupId})
	if err != nil {
		return nil, fmt.Errorf("error selecting group: %w", err)
	}

	group, ok := groupData[0].(*dbmodels.Group)
	if !ok {
		return nil, fmt.Errorf("could not assert group type")
	}

	return group, nil
}
