package userdb

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
)

/*
SelectUser returns a user from the database given a column value and a search value.

Firstly, validates that the column value is present in the database table, so that it can query the database properly.
Then, opens the database, prepares the statement, and finally queries the row before returning the user as a User struct.

Parameters:
- db (*sql.DB): An open database to access and interact with.
- columnValue (string): refers to the column name in the sql table.
- searchValue (string): what to search for in the relevant column.

Returns:
- dbmodels.User: a user struct containing all the details of the queried user.
- error: any error relating to selecting the user.

Errors:
- Returns an error if columnName does not exist within the table by checking the relevant struct.
- Returns an error if the database failed to open.
- Returns an error if preparing the statement failed.
- Returns an error if querying the rows fails.

Examples:
- Used when selecting a user, maybe to retrieve certain details or verify credentials.
*/
func SelectUser(db *sql.DB, columnName, searchValue string) (dbmodels.User, error) {
	var user dbmodels.User

	doesColumnExist, err := dbutils.DoesColumnExist(db, "Users", columnName)
	if !doesColumnExist {
		return user, fmt.Errorf("search column doesn't exist when selecting user: %w", err)
	}

	stm, err := db.Prepare("SELECT * FROM Users WHERE " + columnName + " = ?")
	if err != nil {
		return user, fmt.Errorf("statement preparation failure when selecting user: %w", err)
	}

	err = stm.QueryRow(searchValue).Scan(
		&user.UserId,
		&user.IsLoggedIn,
		&user.Email,
		&user.HashedPassword,
		&user.FirstName,
		&user.LastName,
		&user.DOB,
		&user.AvatarPath,
		&user.DisplayName,
		&user.AboutMe,
		&user.CreationDate,
	)
	if err != nil {
		return user, fmt.Errorf("query row failure when selecting user: %w", err)
	}
	return user, err
}
