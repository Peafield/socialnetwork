package routehandlers

import (
	"encoding/json"
	"net/http"
	"socialnetwork/pkg/controllers/routecontrollers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewGroup(w, r)
		return
	case http.MethodGet:
		GetGroup(w, r)
		return
	case http.MethodPut:
		UpdateGroup(w, r)
		return
	case http.MethodDelete:
		DeleteGroup(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
Implements the POST method for the "/groups" endpoint.
This function will INSERT a new group into the database.
*/
func NewGroup(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	//sanitize group data: check if columns are correct
	err := routecontrollers.InsertGroup(dbutils.DB, userInfo.UserId, groupData.Data)

	if err != nil {
		http.Error(w, "Failed to create a group", http.StatusInternalServerError)
	}

	//redirect or send response?
	w.WriteHeader(http.StatusOK)
}

/*
Implements the GET method for the "/groups" endpoint.
This function will SELECT a number of groups from the database.
*/

// what are the different ways a user can get a group?
func GetGroup(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	//sanitize data
	result, err := routecontrollers.SelectGroup(dbutils.DB, userInfo.UserId, groupData.Data)

	if err != nil {
		http.Error(w, "Failed to GET group", http.StatusInternalServerError)
	}

	//redirect or send Json response?
	//add token to response type, marshal and send back
	response := readwritemodels.WriteData{
		Status: "success",
		Data:   result,
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

/*
Implements the UPDATE method for the "/groups" endpoint.
This function will UPDATE a group in the database if user has the adequate permissions.
*/
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	// check if the user has the necessary permissions to get the group

}

/*
Implements the DELETE method for the "/groups" endpoint.
This function will DELETE a group from the database if the user has the adequate permissions.
*/
func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
