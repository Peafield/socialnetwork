package usercontrollers

import (
	"database/sql"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"strings"
	"time"
)

/*
UpdateUserAccount updates the current users details.

It creates the conditions map to find the record to refer to using the data passed through.  Then
checks to make sure no immutable values are being passed through.  Then the update can take place.

Parameters:
  - db (*sql.DB): an open database with which to interact.
  - userId (string): the current users id.
  - updateUserData (map[string]interface{}): data about the user account to update.

Errors:
  - an immutable property was found in the updateUserData map.
  - failure to update the database
*/
func UpdateUserAccount(db *sql.DB, userId string, updateUserData map[string]interface{}) error {
	var columns []string
	var args []interface{}

	if email, ok := updateUserData["email"].(string); ok {
		columns = append(columns, "email = ?")
		args = append(args, email)
	}
	if firstName, ok := updateUserData["first_name"].(string); ok {
		columns = append(columns, "first_name = ?")
		args = append(args, firstName)
	}
	if lastName, ok := updateUserData["last_name"].(string); ok {
		columns = append(columns, "last_name = ?")
		args = append(args, lastName)
	}
	if dob, ok := updateUserData["data_of_birth"].(time.Time); ok {
		columns = append(columns, "date_of_birth = ?")
		args = append(args, dob)
	}
	if avatarPath, ok := updateUserData["avatar_path"].(string); ok {
		columns = append(columns, "avatar_path = ?")
		args = append(args, avatarPath)
	}
	if displayName, ok := updateUserData["display_name"].(string); ok {
		args = append(args, displayName)
	}
	if aboutMe, ok := updateUserData["about_me"].(string); ok {
		args = append(args, aboutMe)
	}

	query := fmt.Sprintf("UPDATE Users SET %s WHERE user_id = ?", strings.Join(columns, ", "))
	updateUserDetailsStatment, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update post statment: %w", err)
	}
	defer updateUserDetailsStatment.Close()

	args = append(args, userId)

	err = crud.InteractWithDatabase(db, updateUserDetailsStatment, args)
	if err != nil {
		return fmt.Errorf("failed to update post data: %w", err)
	}
	return nil
}
