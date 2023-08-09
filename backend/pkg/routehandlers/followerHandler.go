package routehandlers

import (
	"net/http"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
FollowerHandler is a general HTTP request handler for actions related to followers.
It is designed to process different HTTP methods (GET, POST, PUT, DELETE) and call the corresponding functions for each method.

Based on the HTTP method, it will:

	GET: Retrieve follower/s using GetFollowerHandler(w, r)
	POST: Insert a follower record into the database using NewFollowerHandler(w, r)
	PUT: Update a follower's record using UpdateFollowerHandler(w, r)
	DELETE: Delete a follower's record using DeleteFollowerHandler(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/follower", FollowerHandler)

In this example, FollowerHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/follower" on the server will be processed by FollowerHandler.
*/
func FollowerHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewFollowerHandler(w, r)
		return
	case http.MethodGet:
		//get one follower or get all followers?
		//GetAllFollowers(w, r)
		return
	case http.MethodPut:
		UpdateFollowerHandler(w, r)
		return
	case http.MethodDelete:
		DeleteFollowerHandler(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
NewFollowerHandler is an HTTP handler for creating a new follow record in the web application.
This function extracts user data and follower data from the HTTP request context, then inserts a new follow record into the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the follower data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the follower data from the context.
  - Calls controllers.FollowUser to insert a new follower record into the database. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert a new follower record.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully inserted.
*/
func NewFollowerHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	postFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update post follower data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.FollowUser(dbutils.DB, userData.UserId, postFollowerData.Data)
	if err != nil {
		http.Error(w, "failed to follow user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
UpdateFollowerHandler

COULD BE USED FOR ACCEPTING A FOLLOW REQUEST???
*/
func UpdateFollowerHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	updateFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update follower data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.UpdateFollowStatus(dbutils.DB, userData.UserId, updateFollowerData.Data)
	if err != nil {
		http.Error(w, "failed to update follower status", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
DeleteFollowerHandler is an HTTP handler for deleting a follower from the database.
This function extracts user data and delete follower data from the HTTP request context, then deletes a specific follower
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the delete follower data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the delete follower data from the context.
  - Calls controllers.UnfollowUser to delete a follower. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to delete the follower.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully deleted.
*/
func DeleteFollowerHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	deleteFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read delete follower data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.UnfollowUser(dbutils.DB, userData.UserId, deleteFollowerData.Data)
	if err != nil {
		http.Error(w, "failed to delete follower", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}