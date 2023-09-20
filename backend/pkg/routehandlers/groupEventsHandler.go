package routehandlers

import (
	"encoding/json"
	"log"
	"net/http"
	eventcontrollers "socialnetwork/pkg/controllers/EventControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
	"sort"
)

func GroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		PostGroupEvent(w, r)
		return
	case http.MethodGet:
		GetGroupEvent(w, r)
		return
		// 	case http.MethodPut:
		// 		PutGroupEvent(w, r)
		// 		return
		// 	case http.MethodDelete:
		// 		DeleteGroupEvent(w, r)
		// 		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

// /*
// Implements the POST method within the "/groupsevents" endpoint.
// This function will INSERT a new group into the database.
// */
func PostGroupEvent(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	eventData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	//insert new group event
	err := eventcontrollers.InsertEvent(dbutils.DB, userInfo.UserId, eventData.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to insert event into group", http.StatusInternalServerError)
		return
	}

	response := readwritemodels.WriteData{
		Status: "success",
	}

	jsonReponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonReponse)
	w.WriteHeader(http.StatusOK)
}

// /*
// Implements the GET method within the "/groups" endpoint.
// This function will SELECT a number of groups from the database (for what purpose??).
// */
func GetGroupEvent(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	var response readwritemodels.WriteData
	groupId := r.URL.Query().Get("group_id")

	if groupId != "" {
		result, err := eventcontrollers.SelectAllGroupEvents(dbutils.DB, userInfo.UserId, groupId)
		if err != nil {
			http.Error(w, "Failed to GET Event", http.StatusInternalServerError)
		}

		sort.Slice(result.GroupEvents, func(i, j int) bool {
			return result.GroupEvents[i].EventStartTime.After(result.GroupEvents[j].EventStartTime)
		})

		//add token to response type, marshal and send back
		response = readwritemodels.WriteData{
			Status: "success",
			Data:   result,
		}
	}

	jsonReponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonReponse)
	w.WriteHeader(http.StatusOK)

}

// /*
// Implements the UPDATE method within the "/groups" endpoint.
// This function will UPDATE a group in the database if user has the adequate permissions.
// */
// func PutGroupEvent(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	eventData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	// check whether the map's keys match the expected parameters
// 	expectedParams := []string{"title", "description", "event_start_time", "group_id", "event_id"}
// 	found := helpers.FoundParameters(eventData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in UpdateGroupMember", http.StatusBadRequest)
// 		return
// 	}

// 	//initialize data
// 	groupId, _ := eventData.Data["group_id"].(string)
// 	eventId, _ := eventData.Data["event_id"].(string)

// 	err := groupcontrollers.UpdateEvent(dbutils.DB, userInfo.UserId, groupId, eventId, eventData.Data)
// 	if err != nil {
// 		http.Error(w, "Failed to UPDATE event", http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the DELETE method within the "/groups" endpoint.
// This function will DELETE a group from the database if the user has the adequate permissions.
// */
// func DeleteGroupEvent(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	eventData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"group_id", "event_id"}
// 	found := helpers.FoundParameters(eventData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in UpdateGroupMember", http.StatusBadRequest)
// 		return
// 	}

// 	//initialize data
// 	groupId, _ := eventData.Data["group_id"].(string)
// 	eventId, _ := eventData.Data["event_id"].(string)

// 	err := groupcontrollers.DeleteEvent(dbutils.DB, userInfo.UserId, groupId, eventId)
// 	if err != nil {
// 		http.Error(w, "Failed to DELETE event", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
