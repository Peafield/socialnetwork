package routehandlers

import (
	"net/http"
	postselectedfollowercontrollers "socialnetwork/pkg/controllers/PostSelectedFollowerControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
PostSelectedFollowerHandler is a general HTTP request handler for actions related to post selected followers.
It is designed to process different HTTP methods (GET, POST, PUT, DELETE) and call the corresponding functions for each method.

Based on the HTTP method, it will:

	GET: Retrieve post selected follower/s using GetPostSelectedFollowerHandler(w, r)
	POST: Insert a post selected follower record into the database using NewPostSelectedFollowerHandler(w, r)
	PUT: Update a post selected follower's record using UpdatePostSelectedFollowerHandler(w, r)
	DELETE: Delete a post selected follower's record using DeletePostSelectedFollowerHandler(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/postselectedfollower", PostSelectedFollowerHandler)

In this example, PostSelectedFollowerHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/postselectedfollower" on the server will be processed by PostSelectedFollowerHandler.
*/
func PostSelectedFollowerHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewPostSelectedFollowerHandler(w, r)
		return
	case http.MethodGet:
		//get one follower or get all followers?
		//GetAllFollowers(w, r)
		return
	case http.MethodPut:
		//MIGHT NOT BE NEEDED
		//UpdatePostSelectedFollowerHandler(w, r)
		return
	case http.MethodDelete:
		DeletePostSelectedFollowerHandler(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
NewPostSelectedFollowerHandler is an HTTP handler for creating a new post selected follower record in the web application.
This function extracts user data and post selected follower data from the HTTP request context, then inserts a new post selected follower record into the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the post selected follower data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post selected follower data from the context.
  - Calls controllers.NewPostSelectedFollower to insert a new post selected follower record into the database. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert a new post selected follower record.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully inserted.
*/
func NewPostSelectedFollowerHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newPostSelectedFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read new post selected follower data from context", http.StatusInternalServerError)
		return
	}

	err := postselectedfollowercontrollers.NewPostSelectedFollower(dbutils.DB, userData.UserId, newPostSelectedFollowerData.Data)
	if err != nil {
		http.Error(w, "failed to insert new post selected follower", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
UpdatePostSelectedFollowerHandler
*/
// func UpdatePostSelectedFollowerHandler(w http.ResponseWriter, r *http.Request) {
// 	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	updatePostSelectedFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read update post selected follower data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	err := controllers.UpdatePostSelectedFollower(dbutils.DB, userData.UserId, updatePostSelectedFollowerData.Data)
// 	if err != nil {
// 		http.Error(w, "failed to update post selected follower", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

/*
DeletePostSelectedFollowerHandler is an HTTP handler for deleting a post selected follower from the database.
This function extracts user data and delete post selected follower data from the HTTP request context, then deletes a specific
post selected follower based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the delete post selected follower data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the delete post selected follower data from the context.
  - Calls controllers.DeletePostSelectedFollower to delete a post selected follower. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to delete the post selected follower.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully deleted.
*/
func DeletePostSelectedFollowerHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	deletePostSelectedFollowerData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read delete post selected follower data from context", http.StatusInternalServerError)
		return
	}

	err := postselectedfollowercontrollers.DeletePostSelectedFollower(dbutils.DB, userData.UserId, deletePostSelectedFollowerData.Data)
	if err != nil {
		http.Error(w, "failed to delete post selected follower", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
