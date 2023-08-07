package routehandlers

import "net/http"

func GroupEventAttendeesHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewAttendees(w, r)
		return
	case http.MethodGet:
		GetAttendees(w, r)
		return
	case http.MethodPut:
		UpdateAttendees(w, r)
		return
	case http.MethodDelete:
		DeleteAttendees(w, r)
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
func NewAttendees(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the GET method within the "/groups" endpoint.
This function will SELECT a number of groups from the database (for what purpose??).
*/
func GetAttendees(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the UPDATE method within the "/groups" endpoint.
This function will UPDATE a group in the database if user has the adequate permissions.
*/
func UpdateAttendees(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the DELETE method within the "/groups" endpoint.
This function will DELETE a group from the database if the user has the adequate permissions.
*/
func DeleteAttendees(w http.ResponseWriter, r *http.Request) {

}
