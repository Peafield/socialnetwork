package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"strings"
)

/*
ValidateCredentials checks that a given email or display name match with the given password and returns
the appropriate user if they do.

The formData map is read to retrieve the username/email and the password.  Then, using the given username/email,
a user is selected from the database and the password is compared with the stored hashed password.  If the passwords
are a match, a user is returned, otherwise an error.

Parameters:
  - formData (map[string]interface{}): a map of inputs from the user when trying to sign in.
  - db (*sql.DB): the open database to use.

Returns:
  - *dbmodels.User: if there are no errors in the validation process, a user is returned.
  - error: any errors that might occur.

Errors:
  - if the username/email or password is not a string in the formData map.
  - if there was an error selecting a user from the database (ex. user doesn't exist).
  - if the selected interface wasn't a user when selecting a user.
  - if the inputted password wasn't correct with the stored one.
*/
func ValidateCredentials(formData map[string]interface{}, db *sql.DB) (*dbmodels.User, error) {
	// Email
	username_email, ok := formData["username_email"].(string)
	if !ok {
		return nil, fmt.Errorf("username_email is not a string")
	}

	// Password
	password, ok := formData["password"].(string)
	if !ok {
		return nil, fmt.Errorf("password is not a string")
	}

	//set conditions
	conditions := make(map[string]interface{})
	if strings.Contains(username_email, "@") {
		conditions["email"] = username_email
	} else {
		conditions["display_name"] = username_email
	}

	conditionStatement := dbutils.ConditionStatementConstructor(conditions)

	//get user data as interface
	userData, err := crud.SelectFromDatabase(db, "Users", conditionStatement)
	if err != nil {
		return nil, fmt.Errorf("error selecting user from database: %s", err)
	}

	//assert dbmodels.User type
	user, ok := userData.(*dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("returned database value is not a User struct: %s", err)
	}

	//compare passwords
	err = helpers.CompareHashedPassword(user.HashedPassword, password)
	if err != nil {
		return nil, fmt.Errorf("inputted password incorrect: %s", err)
	}
	return user, nil
}
