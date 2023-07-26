package controllers

import (
	"encoding/json"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
RegisterUser creates a user struct to be inserted into the database.

The function takes the read data and converts any given data into a user struct.
The userId is then added as a UUID and the password is hashed. Once these have been done successfully,
the user data is split into it's values as an interface so they can inserted into the database. Finally,
the userId is returned.

Parameters:
  - formData (readwritemodels.ReadData): any form data received from the route handler.

Returns:
  - string: a user id as a UUID.
  - error: an error is returned if the formData data is not a string, if the data cannot be unmarshalled into
    the user struct, if the UUID or the hashing of the password fails, if the values from the user struct cannot be
    extracted, or if there is an error inserting the user data into the database.
*/
func RegisterUser(formData readwritemodels.ReadData) (*dbmodels.User, error) {
	dataStr, ok := formData.Data.(string)
	if !ok {
		return nil, fmt.Errorf("formData.Data is not a string")
	}

	var user dbmodels.User
	err := json.Unmarshal([]byte(dataStr), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %s", err)
	}

	userId, err := helpers.CreateUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create userId: %s", err)
	}
	user.UserId = userId

	hashedPassedword, err := helpers.HashPassword(user.HashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to hash user's password: %s", err)
	}
	user.HashedPassword = hashedPassedword

	userValues, err := helpers.StructFieldValues(user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user struct values: %s", err)
	}
	err = crud.InsertIntoDatabase(dbutils.DB, dbstatements.InsertUserStmt, userValues)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into database: %s", err)
	}

	return &user, nil
}
