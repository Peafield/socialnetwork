package userdb

import (
	"database/sql"
	"fmt"
	"reflect"
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
func SelectUser(db *sql.DB, table string, conditionStatement string) (interface{}, error) {
	var user dbmodels.User

	stm := "SELECT * FROM Users " + conditionStatement

	result, err := db.Query(stm)
	if err != nil {
		return user, fmt.Errorf("statement query failed when selecting: %w", err)
	}

	for result.Next() {
		result.Scan(StructFieldAddress(user)...)
	}
	if err != nil {
		return user, fmt.Errorf("query row failure when selecting user: %w", err)
	}
	return user, err
}

func StructFieldAddress(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var addresses []interface{}
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		if fieldType.Name != "CreationDate" {
			addresses = append(addresses, v.Field(i).Addr().Interface())
		}

	}
	return addresses
}
