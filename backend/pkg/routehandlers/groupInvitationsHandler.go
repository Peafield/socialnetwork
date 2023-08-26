package routehandlers

// func GroupInvitationsHandler(w http.ResponseWriter, r *http.Request) {
// 	method := r.Method
// 	switch method {
// 	case http.MethodPost:
// 		PostInvitation(w, r)
// 		return
// 	case http.MethodGet:
// 		GetInvitation(w, r)
// 		return
// 	case http.MethodPut:
// 		PutInvitation(w, r)
// 		return
// 	case http.MethodDelete:
// 		DeleteInvitation(w, r)
// 		return
// 	default:
// 		http.Error(w, "invalid method", http.StatusBadRequest)
// 		return
// 	}
// }

// /*
// Implements the POST method within the "/groupsevents" endpoint.
// This function will INSERT a new group into the database.
// */
// func PostInvitation(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	invitationData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"user_id", "group_id"}
// 	found := helpers.FoundParameters(invitationData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in PostInvitation", http.StatusBadRequest)
// 		return
// 	}

// 	groupId, _ := invitationData.Data["group_id"].(string)
// 	userId, _ := invitationData.Data["user_id"].(string)
// 	isCreator := dbutils.IsGroupCreator(dbutils.DB, userId, groupId)

// 	//check if user is invited or invitee
// 	if userInfo.UserId != userId && !isCreator {
// 		//check if user is the creator
// 		http.Error(w, "user has no permissions to POST an invitation on someone else's behalf", http.StatusBadRequest)
// 		return
// 	}

// 	err := groupcontrollers.InsertInvitation(dbutils.DB, groupId, userId, isCreator)
// 	if err != nil {
// 		http.Error(w, "Failed to insert invitation", http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the GET method within the "/groups" endpoint.
// This function will SELECT a number of groups from the database (for what purpose??).
// */
// func GetInvitation(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	invitationData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	result, err := groupcontrollers.SelectInvitation(dbutils.DB, userInfo.UserId, invitationData.Data)

// 	if err != nil {
// 		http.Error(w, "Failed to GET invitation(s)", http.StatusInternalServerError)
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
// Implements the PUT method within the "/invitations" endpoint.
// This function will UPDATE an invitation in the database if user has the adequate permissions.
// */
// func PutInvitation(w http.ResponseWriter, r *http.Request) {
// 	// UPDATE is useless in this case
// }

// /*
// Implements the DELETE method within the "/invitations" endpoint.
// This function will DELETE an invitation from the database if the user has the adequate permissions.
// */
// func DeleteInvitation(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	invitationData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"user_id", "group_id"}
// 	found := helpers.FoundParameters(invitationData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in DeleteAttendees", http.StatusBadRequest)
// 		return
// 	}

// 	//initialize data
// 	groupId, _ := invitationData.Data["group_id"].(string)
// 	userId, _ := invitationData.Data["user_id"].(string)
// 	isCreator := dbutils.IsGroupCreator(dbutils.DB, userId, groupId)

// 	//check if user is invited or invitee
// 	if userInfo.UserId != userId && !isCreator {
// 		//check if user is the creator
// 		http.Error(w, "user has no permissions to DELETE an invitation on someone else's behalf", http.StatusBadRequest)
// 		return
// 	}

// 	err := groupcontrollers.DeleteInvitation(dbutils.DB, userId, groupId)
// 	if err != nil {
// 		http.Error(w, "Failed to DELETE invitation(s)", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
