package crud_test

import (
	"database/sql"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// TO DO: NEED RANDOM INFO GENERATOR
var (
	user = &dbmodels.User{
		UserId:         "821",
		IsLoggedIn:     1,
		Email:          "user@test.com821",
		HashedPassword: "hashed_password821",
		FirstName:      "First821",
		LastName:       "Last821",
		DOB:            time.Now(),
		AvatarPath:     "path/to/avatar821",
		DisplayName:    "User821",
		AboutMe:        "About me821",
	}

	userValues = helpers.StructFieldValues(user)

	post = &dbmodels.Post{
		PostId:           "1",
		GroupId:          "1",
		CreatorId:        "1",
		Title:            "TEST1",
		ImagePath:        "path/to/image",
		Content:          "A whole bunch of nonsense",
		PrivacyLevel:     0,
		AllowedFollowers: "ted, jill, andrew",
		Likes:            100,
		Dislikes:         100000,
	}
	postValues = helpers.StructFieldValues(post)
)

/*TestInsertIntoDatabase tests insertion of data into the database*/
func TestInsertIntoDatabase(t *testing.T) {
	db, err := sql.Open("sqlite3", "../testDB.db")
	if err != nil {
		t.Errorf("failed to open test database: %s", err)
	} else {
		defer db.Close()
	}

	err = dbstatements.InitDBStatements(db)
	if err != nil {
		t.Errorf("failed to initialise database statements: %s", err)
	} else {
		defer dbstatements.CloseDBStatements()
	}

	testCases := []struct {
		Name          string
		Statement     *sql.Stmt
		Values        []interface{}
		ExpectedError bool
	}{
		{
			Name:          "Correctly inserted",
			Statement:     dbstatements.InsertUserStmt,
			Values:        userValues,
			ExpectedError: false,
		},
		{
			Name:          "Incorrect statement",
			Statement:     dbstatements.InsertSessionsStmt,
			Values:        userValues,
			ExpectedError: true,
		},
		{
			Name:          "Row already exists in database",
			Statement:     dbstatements.InsertUserStmt,
			Values:        userValues,
			ExpectedError: true,
		},
		{
			Name:          "Missing values",
			Statement:     dbstatements.InsertPostStmt,
			Values:        postValues[1:],
			ExpectedError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := crud.InsertIntoDatabase(db, tc.Statement, tc.Values)
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
