package controllers

import (
	"fmt"
	"net/http"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
GetUserFromRequestContext returns a user given a specific payload context.

First it retrieves the payload context set when the user has been validated.  Then it creates a map of conditions to generate a
condition statement.  This condition statement is used to select a user from the database.  We then assert that it is in fact
a User type.

Parameters:
- r (*http.Request): the http request.

Returns:
- *dbmodels.User: a User struct populated with the respective user data.
- error

Errors:
- if there was user id in the payload.
- if there was an error selecting a user from the database.
- if the interface{} value returned from the database is not a User type.
*/
func GetUserFromRequestContext(r *http.Request) (*dbmodels.User, error) {
	payloadData, ok := r.Context().Value(middleware.PayloadKey).(readwritemodels.Payload)
	if !ok {
		return nil, fmt.Errorf("failed to read user ID from context")
	}

	conditions := make(map[string]interface{})
	conditions["user_id"] = payloadData.UserId
	conditionStatement := dbutils.UpdateConditionConstructor(conditions)

	userData, err := crud.SelectFromDatabase(dbutils.DB, "Users", conditionStatement)
	if err != nil {
		return nil, fmt.Errorf("error selecting user from database: %s", err)
	}

	user, ok := userData.(dbmodels.User)
	if !ok {
		return nil, fmt.Errorf("returned database value is not a User struct: %s", err)
	}

	return &user, nil
}
