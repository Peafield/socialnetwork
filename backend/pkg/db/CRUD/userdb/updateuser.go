package db

import (
	"fmt"
	"socialnetwork/pkg/models/dbmodels"
)

func UpdateUser(dbOpener dbmodels.DBOpener, user *dbmodels.User) error {
	db, err := dbOpener.Open(dbOpener.GetDriveName(), dbOpener.GetDataSourceName())
	if err != nil {
		return fmt.Errorf("failed to open database when inserting user: %w", err)
	}
	defer db.Close()

	statement, err := db.Prepare(``)
	if err != nil {
		return fmt.Errorf("failed to prepare insert user statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.UserId, user.IsLoggedIn, user.Email, user.HashedPassword, user.FirstName, user.LastName, user.DOB, user.AvatarPath, user.DisplayName, user.AboutMe)
	if err != nil {
		return fmt.Errorf("failed to execute insert user statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when inserting user: %w", err)
	}

	return nil
}
