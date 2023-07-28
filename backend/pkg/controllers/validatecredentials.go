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
 */
func ValidateCredentials(formData map[string]interface{}, db *sql.DB) (*dbmodels.User, error) {
	// Email
	username_email, ok := formData["username_email"].(string)
	if !ok {
		return nil, fmt.Errorf("email is not a string")
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
	user, ok := userData.(dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("returned database value is not a User struct: %s", err)
	}

	//compare passwords
	err = helpers.CompareHashedPassword(user.HashedPassword, password)
	if err != nil {
		return nil, fmt.Errorf("inputted password incorrect: %s", err)
	}
	return &user, nil
}
