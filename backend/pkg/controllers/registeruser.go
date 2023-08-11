package controllers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"strings"
	"time"
)

/*
RegisterUser is a function that creates a new user record to be inserted into the database.

This function takes a map of form data as input, assigns this data to a User struct,
and then inserts this User struct into the database.

The function performs several operations including:
  - Creation of a UUID for the user ID
  - Setting the IsLoggedIn field
  - Extracting and validating form data fields including email, password, first name, last name,
    date of birth, display name, and about me. The password is also hashed for security.
  - Assigning a placeholder avatar path (this may be updated later to provide actual avatar functionality)
  - Extracting the fields of the User struct into a format suitable for database insertion

Parameters:
  - formData (map[string]interface{}): The form data received from the route handler, which contains
    user registration details. It's expected to include fields for email, password, first name,
    last name, date of birth, display name, and about me.

Returns:
  - *dbmodels.User: Pointer to the User struct that was created and inserted into the database.
  - error: An error is returned if there are any issues during the process. This could be due to
    incorrect types in the form data, failure to create the UUID, hashing of the password, extraction
    of User struct fields, or issues with the database insertion.

Errors:
  - Returns an error if the UUID creation fails
  - Returns an error if the password hashing fails
  - Returns an error for each form data field that is not of the expected type
  - Returns an error if there is an issue extracting the User struct fields
  - Returns an error if there is an issue inserting the User struct into the database
*/
func RegisterUser(formData map[string]interface{}, db *sql.DB, statement *sql.Stmt) (*dbmodels.User, error) {
	var args []interface{}

	// UUID
	userId, err := helpers.CreateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create userId: %s", err)
	}
	args = append(args, userId)

	// IsLoggedIn
	args = append(args, 1)

	// email
	if email, ok := formData["email"].(string); ok {
		args = append(args, email)
	}

	// displayName
	if displayName, ok := formData["display_name"].(string); ok {
		args = append(args, displayName)
	}

	//set query statements
	queryStatement := ""
	queryValues := make([]interface{}, 0)
	if strings.Contains(formData["email"].(string), "@") {
		queryStatement = `
			SELECT * FROM Users 
			WHERE email = ?
			`
		queryValues = append(queryValues, formData["email"].(string))
	} else {
		queryStatement = `
			SELECT * FROM Users 
			WHERE display_name = ?
			`
		queryValues = append(queryValues, formData["display_name"].(string))
	}

	//get user data as interface
	users, err := crud.SelectFromDatabase(db, "Users", queryStatement, queryValues)
	if err == nil && len(users) > 0 {
		return nil, fmt.Errorf("user display name or email already in use")
	}

	// Password
	if password, ok := formData["password"].(string); ok {
		hashedPassedword, err := helpers.HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("failed to hash user's password: %s", err)
		}
		args = append(args, hashedPassedword)
	}

	// First Name
	if firstName, ok := formData["first_name"].(string); ok {
		args = append(args, firstName)
	}

	// Last Name
	if lastName, ok := formData["last_name"].(string); ok {
		args = append(args, lastName)
	}

	// DOB
	if dob, ok := formData["dob"].(time.Time); ok {
		args = append(args, dob)
	}

	if avatarPath, ok := formData["avatar_path"].(string); ok {
		args = append(args, avatarPath)
	}

	// About Me
	if aboutMe, ok := formData["about_me"].(string); ok {
		args = append(args, aboutMe)
	}

	err = crud.InteractWithDatabase(db, statement, args)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into database: %s", err)
	}

	//set query statements
	queryStatement = `SELECT * FROM Users WHERE user_id =?`
	queryValues = []interface{}{
		userId,
	}
	userData, err := crud.SelectFromDatabase(db, "Users", queryStatement, queryValues)
	if err != nil {
		return nil, fmt.Errorf("error selecting user from database: %s", err)
	}

	if len(userData) > 1 {
		return nil, fmt.Errorf("found multiple users with same credentials, ???")
	}

	user, ok := userData[0].(*dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("returned database value is not a User struct: %s", err)
	}

	return user, nil
}
