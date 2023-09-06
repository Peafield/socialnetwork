package routehandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/pkg/controllers"
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
	err := controllers.SignOutUser(dbutils.DB, userInfo.UserId, affectedColumns)

	if err != nil {
		log.Println("Error signing out user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	r.Header.Del("Authorization")

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   1,
	}

	jsonReponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReponse)
}
