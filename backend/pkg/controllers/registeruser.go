package controllers

import (
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
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
func RegisterUser(formData map[string]interface{}) (*dbmodels.User, error) {
	var user dbmodels.User

	// UUID
	userId, err := helpers.CreateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create userId: %s", err)
	}
	user.UserId = userId

	// IsLoggedIn
	user.IsLoggedIn = 1

	// Email
	email, ok := formData["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email is not a string")
	}
	user.Email = email

	// Password
	password, ok := formData["password"].(string)
	if !ok {
		return nil, fmt.Errorf("password is not a string")
	}
	hashedPassedword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash user's password: %s", err)
	}
	user.HashedPassword = hashedPassedword

	// First Name
	firstName, ok := formData["first_name"].(string)
	if !ok {
		return nil, fmt.Errorf("first name is not a string")
	}
	user.FirstName = firstName

	// Last Name
	lastName, ok := formData["last_name"].(string)
	if !ok {
		return nil, fmt.Errorf("last name is not a string")
	}
	user.LastName = lastName

	// DOB
	dob, ok := formData["dob"].(time.Time)
	if !ok {
		return nil, fmt.Errorf("dob is not a type of time")
	}
	user.DOB = dob

	// Avatar Path TO DO
	user.AvatarPath = "path/to/image"

	// Display Name
	displayName, ok := formData["display_name"].(string)
	if !ok {
		return nil, fmt.Errorf("display name is not a string")
	}
	user.DisplayName = displayName

	// About Me
	aboutMe, ok := formData["about_me"].(string)
	if !ok {
		return nil, fmt.Errorf("about me is not a string")
	}
	user.AboutMe = aboutMe

	userValues, err := helpers.StructFieldValues(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user struct values: %s", err)
	}
	err = crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertUserStmt, userValues)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into database: %s", err)
	}

	return &user, nil
}
