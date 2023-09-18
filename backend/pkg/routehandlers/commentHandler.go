package routehandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	commentcontrollers "socialnetwork/pkg/controllers/CommentControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"sort"
)

/*
CommentHandler is a general HTTP request handler for actions related to a comments.
It is designed to process different HTTP methods (GET, POST, PUT, DELETE) and call the corresponding functions for each method.

Based on the HTTP method, it will:

	POST: Create a new comment using NewComment(w, r)
	GET: Retrieve a posts comments using PostComments(w, r)
	PUT: Update an existing comment using UpdateComment(w, r)
	DELETE: Delete a comment using DeleteComment(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/comments", CommentHandler)

In this example, CommentHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/comments" on the server will be processed by CommentHandler.
*/
func CommentHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewComment(w, r)
		return
	case http.MethodGet:
		GetComments(w, r)
		return
	case http.MethodPut:
		UpdateComment(w, r)
		return
	case http.MethodDelete:
		DeleteComment(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
NewComment is an HTTP handler for creating a new comment in the web application.
This function extracts user data and comment data from the HTTP request context, then inserts the new comment into the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the comment data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.InsertComment to insert the new comment data into the database. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully inserted.
*/
func NewComment(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newCommentData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read post data from context", http.StatusInternalServerError)
		return
	}

	err := commentcontrollers.InsertComment(dbutils.DB, userData.UserId, newCommentData.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to insert comment data", http.StatusInternalServerError)
		return
	}
	response := readwritemodels.WriteData{
		Status: "success",
		Data:   "",
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
PostComments is an HTTP handler for selecting all comments for a post from the database.
This function extracts comment data from the HTTP request context, then selects all comments related to a post from the database

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the comment data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Calls controllers.SelectPostComments to select all comments related to a post. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it writes the posts to a response and sends an HTTP 200 (OK) status to the client, indicating that the comments were successfully selected.
*/
func GetComments(w http.ResponseWriter, r *http.Request) {
	comments := &dbmodels.Comments{}
	var err error

	postId := r.URL.Query().Get("post_id")

	if postId != "" {
		comments, err = commentcontrollers.SelectPostComments(dbutils.DB, postId)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to select post's comments", http.StatusInternalServerError)
			return
		}
	}

	sort.Slice(comments.Comments, func(i, j int) bool {
		return comments.Comments[i].CommentInfo.CreationDate.After(comments.Comments[j].CommentInfo.CreationDate)
	})

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   comments,
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
UpdateComment is an HTTP handler for updating a comment in the database.
This function extracts user data and update comment data from the HTTP request context, then updates a specific comment
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the update comment data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.UpdateUserComment to update a comment. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully updated.
*/
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	updateCommentData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update comment data from context", http.StatusInternalServerError)
		return
	}

	err := commentcontrollers.UpdateUserComment(dbutils.DB, userData.UserId, updateCommentData.Data)
	if err != nil {
		http.Error(w, "failed to update user comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
DeleteComment is an HTTP handler for deleting a comment from the database.
This function extracts user data and delete comment data from the HTTP request context, then deletes a specific comment
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the comment data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.DeleteUserComment to delete a comment. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the comment was successfully deleted.
*/
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	deleteCommentData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read delete comment data from context", http.StatusInternalServerError)
		return
	}

	err := commentcontrollers.DeleteUserComment(dbutils.DB, userData.UserId, deleteCommentData.Data)
	if err != nil {
		http.Error(w, "failed to delete user comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
