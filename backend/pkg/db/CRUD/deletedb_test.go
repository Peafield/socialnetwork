package crud_test

import (
	"database/sql"
	crud "socialnetwork/pkg/db/CRUD"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// NEED TO CREATE A USER AND THEN DELETE IT
func TestDeleteFromDatabase(t *testing.T) {
	db, err := sql.Open("sqlite3", "../testDB.db")
	if err != nil {
		t.Errorf("failed to open test database: %s", err)
	} else {
		defer db.Close()
	}
	testCases := []struct {
		Name          string
		Table         string
		Conditions    map[string]interface{}
		ExpectedError bool
	}{
		{
			Name:  "Delete Succesfully",
			Table: "Users",
			Conditions: map[string]interface{}{
				"user_id": "5",
			},
			ExpectedError: false,
		},
		{
			Name:  "Row does not exist to delete",
			Table: "Users",
			Conditions: map[string]interface{}{
				"user_id": "1",
			},
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := crud.DeleteFromDatabase(db, tc.Table, tc.Conditions)
			if tc.ExpectedError {
				if err == nil {
					t.Errorf("expected and error but did not get one")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but go one: %s", err)
				}
			}
		})
	}
}
