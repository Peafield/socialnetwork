package routehandlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
UserHandler is a general HTTP request handler for actions related to users.
It is designed to process different HTTP methods (GET, PUT, DELETE) and call the corresponding functions for each method.
The POST method is not accounted for as this is assigned in the sign up handler.

Based on the HTTP method, it will:

	GET: Retrieve user/s using GetUserHandler(w, r)
	PUT: Update a user's account using UpdateUserHandler(w, r)
	DELETE: Delete a user's account using DeleteUserHandler(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/user", UserHandler)

In this example, UserHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/user" on the server will be processed by UserHandler.
*/
func UserHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		GetUserHandler(w, r)
		return
	case http.MethodPut:
		UpdateUserHandler(w, r)
		return
	case http.MethodDelete:
		DeleteUserHandler(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	specificUserId := r.URL.Query().Get("user_id")
	specificUserDisplayName := r.URL.Query().Get("display_name")
	var data interface{}

	if specificUserId == "" && specificUserDisplayName == "" {
		//get all users?
	} else {
		var statement *sql.Stmt
		userDetails := ""

		if specificUserId != "" {
			statement = dbstatements.SelectUserByIDStmt
			userDetails = specificUserId
		} else {
			userDetails = specificUserDisplayName
			statement = dbstatements.SelectUserByDisplayNameStmt
		}

		user, err := usercontrollers.GetUser(dbutils.DB, userData.UserId, statement, userDetails)
		if err != nil {
			http.Error(w, "failed to get specific user data", http.StatusInternalServerError)
		}

		data = user
	}

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   data,
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
UpdateUserHandler is an HTTP handler for updating a user account in the database.
This function extracts user data and update user data from the HTTP request context, then updates the users account.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the update user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the update user data from the context.
  - Calls controllers.UpdateUserAccount to update a user. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to update the user account.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully updated.
*/
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	updateUserData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update user data from context", http.StatusInternalServerError)
		return
	}

	err := usercontrollers.UpdateUserAccount(dbutils.DB, userData.UserId, updateUserData.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update user account", http.StatusInternalServerError)
		return
	}

	user, err := usercontrollers.GetUser(dbutils.DB, userData.UserId, dbstatements.SelectUserByIDStmt, userData.UserId)
	if err != nil {
		http.Error(w, "failed to get specific user data", http.StatusInternalServerError)
	}

	data = user

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   data,
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
DeleteUserHandler is an HTTP handler for deleting a user account from the database.
This function extracts user data and delete user data from the HTTP request context, then deletes a specific user
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the delete user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the delete user data from the context.
  - Calls controllers.DeleteUserAccount to delete a user. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to delete the user account.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully deleted.
*/
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	// USERS SHOULD HAVE TO RETYPE THEIR CREDENTIALS TO DELETE THEIR OWN ACCOUNT
	deleteUserData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read delete user data from context", http.StatusInternalServerError)
		return
	}

	err := usercontrollers.DeleteUserAccount(dbutils.DB, userData.UserId, deleteUserData.Data)
	if err != nil {
		http.Error(w, "failed to delete user account", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
