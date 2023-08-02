package routehandlers

import "net/http"

func GroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewGroupEvent(w, r)
		return
	case http.MethodGet:
		GetGroupEvent(w, r)
		return
	case http.MethodPut:
		UpdateGroupEvent(w, r)
		return
	case http.MethodDelete:
		DeleteGroupEvent(w, r)
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
func NewGroupEvent(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the GET method within the "/groups" endpoint.
This function will SELECT a number of groups from the database (for what purpose??).
*/
func GetGroupEvent(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the UPDATE method within the "/groups" endpoint.
This function will UPDATE a group in the database if user has the adequate permissions.
*/
func UpdateGroupEvent(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the DELETE method within the "/groups" endpoint.
This function will DELETE a group from the database if the user has the adequate permissions.
*/
func DeleteGroupEvent(w http.ResponseWriter, r *http.Request) {

}
