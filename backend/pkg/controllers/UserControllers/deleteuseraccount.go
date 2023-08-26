package usercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
)

/*
DeleteUserAccount is used to delete a user's account.

In order to delete ones account, you have to retype your credentials and be currently logged in (user id from context).
This function takes the credentials and compares them to the ones in the database.  If they match and the password matches,
then it allows the deletion of the account from the database.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - deleteUserData (map[string]interface{}): data about the user to delete.

Errors:
  - assertion failure of data types.
  - error when selecting the user.
  - error when any of the credentials don't match.
  - failure to delete the follower record from the database.
*/
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
	queryValues := make([]interface{}, 0)
	queryValues = append(queryValues, email)
	queryValues = append(queryValues, username)
	queryValues = append(queryValues, userId)

	//select user from inputted information
	userData, err := crud.SelectFromDatabase(db, "Users", dbstatements.SelectUserByEmailAndDisplayNameAndUserIdStmt, queryValues)
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

	args := []interface{}{}
	args = append(args, userId)

	err = crud.InteractWithDatabase(db, dbstatements.DeleteUserAccountStmt, args)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %w", err)
	}
	return nil
}
