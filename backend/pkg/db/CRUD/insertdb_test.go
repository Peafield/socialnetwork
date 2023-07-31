package crud_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"strconv"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// TO DO: NEED RANDOM INFO GENERATOR
var (
	user = &dbmodels.User{
		UserId:         strconv.Itoa(rand.Intn(time.Now().Nanosecond())),
		IsLoggedIn:     1,
		Email:          "user" + strconv.Itoa(rand.Intn(time.Now().Nanosecond())) + "@test.com",
		HashedPassword: "hashed_password821",
		FirstName:      "First821",
		LastName:       "Last821",
		DOB:            time.Now(),
		AvatarPath:     "path/to/avatar821",
		DisplayName:    "User" + strconv.Itoa(rand.Intn(time.Now().Nanosecond())),
		AboutMe:        "About me821",
	}

	userValues, _ = helpers.StructFieldValues(user)

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
	postValues, _ = helpers.StructFieldValues(post)

	chat = &dbmodels.Chat{
		ChatId:       "ragrg4r4524",
		SenderId:     "gewtygibo5",
		ReceiverId:   "usnb08t79bwv75v",
		CreationDate: time.Now(),
	}

	chatValues, _ = helpers.StructFieldValues(chat)
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
			fmt.Println(tc.Values)
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
