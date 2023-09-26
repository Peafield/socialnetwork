package crud

import (
	"database/sql"
	"fmt"
	"log"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/helpers"
)

/*
SelectFromDatabase selects from a given table using a pre-made condition statement.

First, it decides the struct type to use to store the data which is defined by the table parameter passed through.
Then uses the condition statement and table string to create a new statement that is prepared, then queried.
The result is then scanned into the struct defined at the beginning and then returned.

Parameters:
- db (*sql.DB): an open database with which to interact.
- table (string): the name of a table within the database in which to search.
- conditionStatement (string): the conditions on which to search. Ex: "WHERE user_id = 3 AND email = "example@gmail.com"".

Returns:
- interface{}: a relevant struct containing the data from the database.
- error: if there was a problem selecting something.

Errors:
- no struct relating to the table or the table does not exist.
- failure to prepare select statement.
- failure to execute select query.
- failure to scan data into object.

Example:
- SelectFromDatabase(db, "Users",  "WHERE user_id = 3 AND email = "example@gmail.com"") will return a user
- SelectFromDatabase(db, "Posts",  "WHERE post_id = 25 AND group_id = 4") will return a post
*/
func SelectFromDatabase(db *sql.DB, table string, queryStatement *sql.Stmt, queryValues []interface{}) ([]interface{}, error) {
	object, err := helpers.DecideStructType(table)
	if err != nil {
		return nil, fmt.Errorf("no valid struct with table, or not a valid table, when selecting from database. err: %w", err)
	}

	objectArray := make([]interface{}, 0)

	log.Println(queryStatement)
	log.Println(queryValues)

	result, err := queryStatement.Query(queryValues...)
	if err != nil {
		return objectArray, fmt.Errorf("failed to execute select query: %w", err)
	}
	defer result.Close()

	found := false
	for result.Next() {
		objAddresses, oErr := helpers.StructFieldAddress(object)
		if oErr != nil {
			return objectArray, fmt.Errorf("failed to get struct addresses when selecting from database: %w", err)
		}
		err := result.Scan(objAddresses...)
		if err != nil {
			return objectArray, fmt.Errorf("failed to scan data: %w", err)
		}
		objectArray = append(objectArray, object)
		found = true
		// Declare a new object to avoid rewriting original object
		object, err = helpers.DecideStructType(table)
		if err != nil {
			return nil, fmt.Errorf("no valid struct with table, or not a valid table, when selecting from database. err: %w", err)
		}
	}
	if !found {
		err = errorhandling.ErrNoResultsFound
	}
	return objectArray, err
}
