package routehandlers

import "net/http"

func GroupMembersHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewGroupMember(w, r)
		return
	case http.MethodGet:
		GetGroupMember(w, r)
		return
	case http.MethodPut:
		UpdateGroupMember(w, r)
		return
	case http.MethodDelete:
		DeleteGroupMember(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
Implements the POST method within the "/groups" endpoint.
This function will INSERT a new group into the database.
*/
func NewGroupMember(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the GET method within the "/groups" endpoint.
This function will SELECT a number of groups from the database (for what purpose??).
*/
func GetGroupMember(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the UPDATE method within the "/groups" endpoint.
This function will UPDATE a group in the database if user has the adequate permissions.
*/
func UpdateGroupMember(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the DELETE method within the "/groups" endpoint.
This function will DELETE a group from the database if the user has the adequate permissions.
*/
func DeleteGroupMember(w http.ResponseWriter, r *http.Request) {

}
