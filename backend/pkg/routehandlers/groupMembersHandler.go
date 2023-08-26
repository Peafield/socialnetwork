package routehandlers

// func GroupMembersHandler(w http.ResponseWriter, r *http.Request) {
// 	method := r.Method
// 	switch method {
// 	case http.MethodPost:
// 		NewGroupMember(w, r)
// 		return
// 	case http.MethodGet:
// 		GetGroupMember(w, r)
// 		return
// 	case http.MethodPut:
// 		UpdateGroupMember(w, r)
// 		return
// 	case http.MethodDelete:
// 		DeleteGroupMember(w, r)
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
// func NewGroupMember(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"group_id"}
// 	found := helpers.FoundParameters(groupData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in NewGroupMember", http.StatusBadRequest)
// 		return
// 	}

// 	groupId := groupData.Data["group_id"].(string)

// 	err := groupcontrollers.InsertMember(dbutils.DB, userInfo.UserId, groupId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the GET method within the "/groups" endpoint.
// This function will SELECT a number of groups from the database (for what purpose??).
// */
// func GetGroupMember(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//sanitize data
// 	result, err := groupcontrollers.SelectMember(dbutils.DB, userInfo.UserId, groupData.Data)

// 	if err != nil {
// 		http.Error(w, "Failed to GET Member", http.StatusInternalServerError)
// 	}

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
// 	w.Write(jsonReponse)
// 	w.WriteHeader(http.StatusOK)

// }

// /*
// Implements the UPDATE method within the "/groups" endpoint.
// This function will UPDATE a group in the database if user has the adequate permissions.
// */
// func UpdateGroupMember(w http.ResponseWriter, r *http.Request) {
// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"group_id", "user_id", "permission_level"}
// 	found := helpers.FoundParameters(groupData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in UpdateGroupMember", http.StatusBadRequest)
// 		return
// 	}

// 	err := groupcontrollers.UpdateMember(dbutils.DB, userInfo.UserId, groupData.Data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// /*
// Implements the DELETE method within the "/groups" endpoint.
// This function will DELETE a group from the database if the user has the adequate permissions.
// */
// func DeleteGroupMember(w http.ResponseWriter, r *http.Request) {

// 	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	groupData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	//check whether the map's keys match the expected parameters
// 	expectedParams := []string{"group_id", "user_id"}
// 	found := helpers.FoundParameters(groupData.Data, expectedParams)
// 	if !found {
// 		http.Error(w, "expected parameters not found in DeleteGroupMember", http.StatusBadRequest)
// 		return
// 	}

// 	groupId := groupData.Data["group_id"].(string)
// 	memberId := groupData.Data["user_id"].(string)

// 	err := groupcontrollers.DeleteMember(dbutils.DB, userInfo.UserId, memberId, groupId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)

// }
