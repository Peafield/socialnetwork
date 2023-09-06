package usercontrollers

import (
	"database/sql"
	"fmt"
	"os"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
)

func GetUser(db *sql.DB, userId string, statement *sql.Stmt, specificUserDetails string) (*dbmodels.UserProfileData, error) {
	//should add a way of only displaying certain information based on follow status or privacy level?

	specificUserData, err := crud.SelectFromDatabase(db, "Users", statement, []interface{}{specificUserDetails})
	if err != nil {
		return nil, fmt.Errorf("failed to select user from database: %w", err)
	}

	userData, ok := specificUserData[0].(*dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("cannot assert user type")
	}

	//hide hashed password
	userData.HashedPassword = ""

	user := &dbmodels.UserProfileData{}
	user.UserInfo = *userData
	user.ProfilePic, err = os.ReadFile(userData.AvatarPath)

	return user, nil
}

func getUserPrivate(db *sql.DB, userId string, statement *sql.Stmt, specificUserDetails string) (*dbmodels.UserProfileData, error) {
	//should add a way of only displaying certain information based on follow status or privacy level?

	specificUserData, err := crud.SelectFromDatabase(db, "Users", statement, []interface{}{specificUserDetails})
	if err != nil {
		return nil, fmt.Errorf("failed to select user from database: %w", err)
	}

	userData, ok := specificUserData[0].(*dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("cannot assert user type")
	}

	user := &dbmodels.UserProfileData{}
	user.UserInfo = *userData
	user.ProfilePic, err = os.ReadFile(userData.AvatarPath)

	return user, nil
}
