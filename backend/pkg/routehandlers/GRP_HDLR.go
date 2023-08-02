package routehandlers

import "net/http"

/*we may need to create an interface for GET, PUT, POST and DElETE*/

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

}

/*
Implements the GET method for the "/groups" endpoint.
This function will SELECT a number of groups from the database (for what purpose??).
*/
func GetGroup(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the UPDATE method for the "/groups" endpoint.
This function will UPDATE a group in the database if user has the adequate permissions.
*/
func UpdateGroup(w http.ResponseWriter, r *http.Request) {

}

/*
Implements the DELETE method for the "/groups" endpoint.
This function will DELETE a group from the database if the user has the adequate permissions.
*/
func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
