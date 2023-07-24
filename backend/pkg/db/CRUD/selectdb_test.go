package crud_test

import (
	"database/sql"
	"fmt"
	"reflect"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"testing"
	"time"
)

const TEST_USER_ID = "yfefde47"
const TEST_POST_ID = "4fdddry7"

var (
	user1 = &dbmodels.User{
		UserId:         TEST_USER_ID,
		IsLoggedIn:     1,
		Email:          "user" + TEST_USER_ID + "@test.com",
		HashedPassword: "hashed_password821",
		FirstName:      "First821",
		LastName:       "Last821",
		DOB:            helpers.NormalizeTime(time.Date(2000, 6, 24, 2, 34, 1, 50, time.Local)),
		AvatarPath:     "path/to/avatar821",
		DisplayName:    "User" + TEST_USER_ID,
		AboutMe:        "About me821",
	}

	user1Values, _ = helpers.StructFieldValues(user1)

	user2 = &dbmodels.User{
		UserId:         TEST_USER_ID + "2",
		IsLoggedIn:     1,
		Email:          "user" + TEST_USER_ID + "2" + "@test.com",
		HashedPassword: "hashed_password821",
		FirstName:      "First821",
		LastName:       "Last821",
		DOB:            helpers.NormalizeTime(time.Date(1998, 3, 12, 2, 34, 1, 50, time.Local)),
		AvatarPath:     "path/to/avatar821",
		DisplayName:    "User" + TEST_USER_ID + "2",
		AboutMe:        "About me821",
	}

	user2Values, _ = helpers.StructFieldValues(user2)

	post1 = &dbmodels.Post{
		PostId:           TEST_POST_ID,
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
	post1Values, _ = helpers.StructFieldValues(post1)

	post2 = &dbmodels.Post{
		PostId:           TEST_POST_ID + "2",
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
	post2Values, _ = helpers.StructFieldValues(post2)

	chat1 = &dbmodels.Chat{
		ChatId:       "ragrg4r4524",
		SenderId:     "gewtygibo5",
		ReceiverId:   "usnb08t79bwv75v",
		CreationDate: time.Now(),
	}

	chat1Values, _ = helpers.StructFieldValues(chat1)
)

func TestSelectFromDatabase(t *testing.T) {
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
		Name            string
		Table           string
		Insert          bool
		Conditions      map[string]interface{}
		Value           interface{}
		InsertStatement *sql.Stmt
		InsertValues    []interface{}
		IsSame          bool
		ExpectedError   bool
	}{
		{
			Name:   "Found User Correctly Single Condition",
			Table:  "Users",
			Insert: true,
			Conditions: map[string]interface{}{
				"user_id": user1.UserId,
			},
			Value:           user1,
			InsertStatement: dbstatements.InsertUserStmt,
			InsertValues:    user1Values,
			IsSame:          true,
			ExpectedError:   false,
		},
		{
			Name:   "Found User Correctly Multiple Conditions",
			Table:  "Users",
			Insert: true,
			Conditions: map[string]interface{}{
				"email":           user2.Email,
				"hashed_password": user2.HashedPassword,
			},
			Value:           user2,
			InsertStatement: dbstatements.InsertUserStmt,
			InsertValues:    user2Values,
			IsSame:          true,
			ExpectedError:   false,
		},
		{
			Name:   "Found Post Correctly Single Condition",
			Table:  "Posts",
			Insert: true,
			Conditions: map[string]interface{}{
				"post_id": post1.PostId,
			},
			Value:           post1,
			InsertStatement: dbstatements.InsertPostStmt,
			InsertValues:    post1Values,
			IsSame:          true,
			ExpectedError:   false,
		},
		{
			Name:   "Found Post Correctly Multiple Conditions",
			Table:  "Posts",
			Insert: true,
			Conditions: map[string]interface{}{
				"group_id":   post2.GroupId,
				"creator_id": post2.CreatorId,
			},
			Value:           post2,
			InsertStatement: dbstatements.InsertPostStmt,
			InsertValues:    post2Values,
			IsSame:          true,
			ExpectedError:   false,
		},
		{
			Name:   "Didn't Find Post Correctly Single Condition",
			Table:  "Posts",
			Insert: false,
			Conditions: map[string]interface{}{
				"post_id": "xxx",
			},
			Value:           post2,
			InsertStatement: dbstatements.InsertPostStmt,
			InsertValues:    post2Values,
			IsSame:          false,
			ExpectedError:   true,
		},
		{
			Name:   "Incorrect Table",
			Table:  "Chatss",
			Insert: false,
			Conditions: map[string]interface{}{
				"chat_id": "1",
			},
			Value:           nil,
			InsertStatement: dbstatements.InsertChatsStmt,
			InsertValues:    chat1Values,
			IsSame:          false,
			ExpectedError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Insert {
				err := crud.InsertIntoDatabase(db, tc.InsertStatement, tc.InsertValues)
				fmt.Println(err)
			}

			conditionStatement := dbutils.UpdateConditionConstructor(tc.Conditions)
			obj, err := crud.SelectFromDatabase(db, tc.Table, conditionStatement)
			fmt.Println(err)

			objValues, _ := helpers.StructFieldValues(obj)
			tcValues, _ := helpers.StructFieldValues(tc.Value)

			if tc.IsSame && !reflect.DeepEqual(objValues, tcValues) {
				fmt.Println("returned: ", objValues)
				fmt.Println("expected: ", tcValues)
				t.Errorf("expected value is not equal to actual value")
			} else if !tc.IsSame && !reflect.DeepEqual(objValues, tcValues) {
				err = fmt.Errorf("did not return anything")
			}
			if tc.ExpectedError {
				if err == nil {
					t.Errorf("expected an error but did not get one")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but go one: %s", err)
				}
			}
		})
	}
}
