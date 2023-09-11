package usercontrollers

import (
	"database/sql"
	"fmt"
	imagecontrollers "socialnetwork/pkg/controllers/ImageControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
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
	var args []interface{}

	user, err := getUserPrivate(db, userId, dbstatements.SelectUserByIDStmt, userId)
	if err != nil {
		return fmt.Errorf("error getting user")
	}

	err = validatePassword(user, updateUserData)
	if err != nil {
		return err
	}

	userDataKeys := []string{
		"email",
		"display_name",
		"new_password",
		"first_name",
		"last_name",
		"dob",
		"avatar_path",
		"about_me",
		"is_private",
	}
	userDataOldValues := []interface{}{
		user.UserInfo.Email,
		user.UserInfo.DisplayName,
		user.UserInfo.HashedPassword,
		user.UserInfo.FirstName,
		user.UserInfo.LastName,
		user.UserInfo.DOB,
		user.UserInfo.AvatarPath,
		user.UserInfo.AboutMe,
		user.UserInfo.IsPrivate,
	}

	for i, v := range userDataKeys {
		err := appendOldOrNew(v, updateUserData, userDataOldValues[i], &args)
		if err != nil {
			return fmt.Errorf("failed to append data to args: %w", err)
		}
	}

	args = append(args, userId)

	err = crud.InteractWithDatabase(db, dbstatements.UpdateUserAccountStmt, args)
	if err != nil {
		return fmt.Errorf("failed to update post data: %w", err)
	}
	return nil
}

func appendOldOrNew(valueStr string, updateUserData map[string]interface{}, oldValue interface{}, args *[]interface{}) error {
	if valueStr == "is_private" {
		if value, ok := updateUserData[valueStr].(float64); ok {
			*args = append(*args, int(value))
		} else {
			*args = append(*args, oldValue)
		}
		return nil
	}

	if value, ok := updateUserData[valueStr].(string); ok && value != "" {
		if valueStr == "new_password" {
			hashedPassword, err := helpers.HashPassword(value)
			if err != nil {
				return fmt.Errorf("failed to hash user's password: %s", err)
			}
			value = hashedPassword
		}
		if valueStr == "avatar_path" {
			avatarFilePath, err := imagecontrollers.DecodeImage(value)
			if err != nil {
				return fmt.Errorf("error decoding image: %w", err)
			}
			value = avatarFilePath
		}
		if valueStr == "dob" {
			value, _, _ = strings.Cut(value, "T00:00:00Z")
			formattedDOB, err := time.Parse("2006-01-02", value)
			if err != nil {
				return fmt.Errorf("DOB string can't be parsed into time.Time: %w", err)
			}
			*args = append(*args, formattedDOB)
		} else {
			*args = append(*args, value)
		}
	} else {
		*args = append(*args, oldValue)
	}
	return nil
}

func validatePassword(user *dbmodels.UserProfileData, updateUserData map[string]interface{}) error {
	if oldPassword, ok := updateUserData["old_password"].(string); ok {
		err := helpers.CompareHashedPassword(user.UserInfo.HashedPassword, oldPassword)
		if err != nil {
			return fmt.Errorf("inputted password incorrect: %s", err)
		}
	} else {
		return fmt.Errorf("could not find and validate password")
	}
	return nil
}
