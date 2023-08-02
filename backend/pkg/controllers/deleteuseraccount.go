package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

func DeleteUserAccount(db *sql.DB, userId string, deleteUserData map[string]interface{}) error {
	username, ok := deleteUserData["display_name"].(string)
	if !ok {
		return fmt.Errorf("username is not a string")
	}

	email, ok := deleteUserData["email"].(string)
	if !ok {
		return fmt.Errorf("email is not a string")
	}

	password, ok := deleteUserData["password"].(string)
	if !ok {
		return fmt.Errorf("password is not a string")
	}

	//set query statements
	queryStatement := ""
	queryValues := make([]interface{}, 0)
	queryStatement = `
	SELECT * FROM Users 
	WHERE email = ?
	AND display_name = ?
	`
	queryValues = append(queryValues, email)
	queryValues = append(queryValues, username)

	//select user from inputted information
	userData, err := crud.SelectFromDatabase(db, "Users", queryStatement, queryValues)
	if err != nil {
		return fmt.Errorf("error selecting user when trying to delete account, err: %s", err)
	}

	//assert type
	user, ok := userData[0].(*dbmodels.User)
	if !ok {
		return fmt.Errorf("user is not a User struct")
	}

	//compare user id
	if user.UserId != userId {
		return fmt.Errorf("inputted user id does not match when deleting user")
	}

	//compare password
	err = helpers.CompareHashedPassword(user.HashedPassword, password)
	if err != nil {
		return fmt.Errorf("error comparing passwords when deleting user, err: %s", err)
	}

	//delete user
	err = crud.DeleteFromDatabase(db, "Users", "user_id", userId)
	if err != nil {
		return fmt.Errorf("error deleting user from database, err: %s", err)
	}

	return nil
}
