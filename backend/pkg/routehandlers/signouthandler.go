package routehandlers

import (
	"net/http"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	conditions := make(map[string]interface{})
	affectedColums := make(map[string]interface{})
	conditions["user_id"] = userInfo.UserId
	affectedColums["is_logged_in"] = 0
	crud.UpdateDatabaseRow(dbutils.DB, "Users", conditions, affectedColums)

	r.Header.Del("Authorization")

	w.WriteHeader(http.StatusOK)
}
