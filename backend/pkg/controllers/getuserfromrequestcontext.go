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
