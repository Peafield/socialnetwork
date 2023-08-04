package routehandlers

import (
	"log"
	"net/http"
	"socialnetwork/pkg/controllers/routecontrollers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	affectedColumns := make(map[string]interface{})
	affectedColumns["is_logged_in"] = 0
	err := routecontrollers.SignOutUser(dbutils.DB, userInfo.UserId, affectedColumns)

	if err != nil {
		log.Println("Error signing out user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	r.Header.Del("Authorization")

	w.WriteHeader(http.StatusOK)
}
