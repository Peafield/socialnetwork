package routehandlers

// func GroupEventAttendeesHandler(w http.ResponseWriter, r *http.Request) {
// 	method := r.Method
// 	switch method {
// 	case http.MethodPost:
// 		PostAttendees(w, r)
// 		return
// 	case http.MethodGet:
// 		GetAttendees(w, r)
// 		return
// 	case http.MethodPut:
// 		PutAttendees(w, r)
// 		return
// 	case http.MethodDelete:
// 		DeleteAttendees(w, r)
// 		return
// 	default:
// 		http.Error(w, "invalid method", http.StatusBadRequest)
// 		return
// 	}
// }

// /*
// Implements the POST method within the "/groups" endpoint.
// This function will INSERT a new group into the database.
// */
// func PostAttendees(w http.ResponseWriter, r *http.Request) {
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
// 	expectedParams := []string{"attendee_status", "event_id", "group_id"}
// 	found := helpers.FoundParameters(eventData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in PostAttendees", http.StatusBadRequest)
// 		return
// 	}

// 	eventId, _ := eventData.Data["event_id"].(string)

// 	err := groupcontrollers.InsertAttendee(dbutils.DB, userInfo.UserId, eventId, eventData.Data)
// 	if err != nil {
// 		http.Error(w, "Failed to insert attendee", http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the GET method within the "/groups" endpoint.
// This function will SELECT a number of groups from the database (for what purpose??).
// */
// func GetAttendees(w http.ResponseWriter, r *http.Request) {
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

// 	eventId, ok := eventData.Data["event_id"].(string)
// 	if !ok {
// 		http.Error(w, "failed to read attendee id", http.StatusBadRequest)
// 	}

// 	result, err := groupcontrollers.SelectAttendee(dbutils.DB, userInfo.UserId, eventId, eventData.Data)

// 	if err != nil {
// 		http.Error(w, "Failed to GET attendee", http.StatusInternalServerError)
// 	}

// 	//redirect or send Json response?
// 	//add token to response type, marshal and send back
// 	response := readwritemodels.WriteData{
// 		Status: "success",
// 		Data:   result,
// 	}

// 	jsonReponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonReponse)
// }

// /*
// Implements the UPDATE method within the "/groups" endpoint.
// This function will UPDATE a group in the database if user has the adequate permissions.
// */
// func PutAttendees(w http.ResponseWriter, r *http.Request) {
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
// 	expectedParams := []string{"attendee_id", "event_id", "attending_status", "group_id"}
// 	found := helpers.FoundParameters(eventData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in PutAttendees", http.StatusBadRequest)
// 		return
// 	}

// 	//initialize data
// 	groupId, _ := eventData.Data["group_id"].(string)
// 	eventId, _ := eventData.Data["event_id"].(string)
// 	attendeeId, _ := eventData.Data["attendee_id"].(string)
// 	attendingStatus, _ := eventData.Data["attending_Status"].(bool)

// 	err := groupcontrollers.UpdateAttendee(dbutils.DB, userInfo.UserId, attendeeId, groupId, eventId, attendingStatus)
// 	if err != nil {
// 		http.Error(w, "Failed to UPDATE attendees", http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the DELETE method within the "/groups" endpoint.
// This function will DELETE a group from the database if the user has the adequate permissions.
// */
// func DeleteAttendees(w http.ResponseWriter, r *http.Request) {
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
// 	expectedParams := []string{"attendee_id", "group_id", "event_id"}
// 	found := helpers.FoundParameters(eventData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in DeleteAttendees", http.StatusBadRequest)
// 		return
// 	}

// 	//initialize data
// 	attendeeId, _ := eventData.Data["attendee_id"].(string)
// 	groupId, _ := eventData.Data["group_id"].(string)
// 	eventId, _ := eventData.Data["event_id"].(string)

// 	err := groupcontrollers.DeleteAttendee(dbutils.DB, userInfo.UserId, attendeeId, eventId, groupId)
// 	if err != nil {
// 		http.Error(w, "Failed to DELETE attendees", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
