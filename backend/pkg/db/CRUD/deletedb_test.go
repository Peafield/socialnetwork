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
		Column        string
		Value         string
		ExpectedError bool
	}{
		{
			Name:          "Delete Succesfully",
			Table:         "Users",
			Column:        "user_id",
			Value:         "5",
			ExpectedError: false,
		},
		{
			Name:          "Row does not exist to delete",
			Table:         "Users",
			Column:        "user_id",
			Value:         "1",
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := crud.DeleteFromDatabase(db, tc.Table, tc.Column, tc.Value)
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
