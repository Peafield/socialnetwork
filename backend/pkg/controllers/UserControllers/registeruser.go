package usercontrollers

import (
	"database/sql"
	"fmt"
	imagecontrollers "socialnetwork/pkg/controllers/ImageControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
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
func RegisterUser(db *sql.DB, formData map[string]interface{}) (*dbmodels.User, error) {
	var args []interface{}

	// UUID
	userId, err := helpers.CreateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create userId: %s", err)
	}
	args = append(args, userId)

	// Set IsLoggedIn Status
	args = append(args, 1)

	formDataValues, err := validateAndSortIncomingFormData(formData, db)
	if err != nil {
		return nil, err
	}
	args = append(args, formDataValues...)

	//insert into db
	err = crud.InteractWithDatabase(db, dbstatements.InsertUserStmt, args)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into database: %s", err)
	}

	user, err := selectRegisteredUser(db, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func validateAndSortIncomingFormData(formData map[string]interface{}, db *sql.DB) ([]interface{}, error) {
	var args []interface{}

	emailAndDisplayName, err := validateEmailAndDisplayName(formData, db)
	if err != nil {
		return nil, err
	}
	args = append(args, emailAndDisplayName...)

	otherRequiredFields, err := validateOtherRequiredFields(formData)
	if err != nil {
		return nil, err
	}
	args = append(args, otherRequiredFields...)

	remainingUnrequiredFields, err := validateAnyUnrequiredFields(formData)
	if err != nil {
		return nil, err
	}
	args = append(args, remainingUnrequiredFields...)

	return args, nil
}

func validateEmailAndDisplayName(formData map[string]interface{}, db *sql.DB) ([]interface{}, error) {
	var args []interface{}

	// email
	if email, ok := formData["email"].(string); ok && strings.Contains(email, "@") {
		args = append(args, email)
	} else {
		return nil, fmt.Errorf("email is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	// displayName
	if displayName, ok := formData["display_name"].(string); ok {
		args = append(args, displayName)
	} else {
		return nil, fmt.Errorf("display name is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	//get user data as interface
	users, err := crud.SelectFromDatabase(db, "Users", dbstatements.SelectUserByIDOrDisplayNameStmt, args)
	if err == nil && len(users) > 0 {
		return nil, errorhandling.ErrUserExists
	}

	return args, nil
}

func validateOtherRequiredFields(formData map[string]interface{}) ([]interface{}, error) {
	var args []interface{}

	// Password
	if password, ok := formData["password"].(string); ok {
		hashedPassedword, err := helpers.HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("failed to hash user's password: %s", err)
		}
		args = append(args, hashedPassedword)
	} else {
		return nil, fmt.Errorf("password is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	// First Name
	if firstName, ok := formData["first_name"].(string); ok {
		args = append(args, firstName)
	} else {
		return nil, fmt.Errorf("firstName is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	// Last Name
	if lastName, ok := formData["last_name"].(string); ok {
		args = append(args, lastName)
	} else {
		return nil, fmt.Errorf("lastName is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	// DOB
	if dob, ok := formData["dob"].(string); ok {
		formattedDOB, err := time.Parse("2006-01-02", dob)
		if err != nil {
			return nil, fmt.Errorf("DOB string can't be parsed into time.Time: %w", err)
		}
		args = append(args, formattedDOB)
	} else {
		return nil, fmt.Errorf("DOB is not a string, or is not in a valid format, or doesn't exist when it should")
	}

	return args, nil
}

func validateAnyUnrequiredFields(formData map[string]interface{}) ([]interface{}, error) {
	var args []interface{}

	if avatarPath, ok := formData["avatar_path"].(string); ok {
		avatarFilePath, err := imagecontrollers.DecodeImage(avatarPath)
		if err != nil {
			return nil, fmt.Errorf("error decoding image: %w", err)
		}
		args = append(args, avatarFilePath)
	} else {
		args = append(args, "")
	}

	// About Me
	if aboutMe, ok := formData["about_me"].(string); ok {
		args = append(args, aboutMe)
	} else {
		args = append(args, "")
	}

	return args, nil
}

func selectRegisteredUser(db *sql.DB, userId string) (*dbmodels.User, error) {
	//set query values
	queryValues := []interface{}{
		userId,
	}
	userData, err := crud.SelectFromDatabase(db, "Users", dbstatements.SelectUserByIDStmt, queryValues)
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
