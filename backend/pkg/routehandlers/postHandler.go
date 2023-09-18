package routehandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	postcontrollers "socialnetwork/pkg/controllers/PostControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"sort"
)

/*
PostHandler function is a general HTTP request handler for actions related to a post in a web application.
It is designed to process different HTTP methods (GET, POST, PUT, DELETE) and call the corresponding functions for each method.

Based on the HTTP method, it will:

	POST: Create a new post using NewPost(w, r)
	GET: Retrieve a users posts using UserPosts(w, r)
	PUT: Update an existing post using UpdatePost(w, r)
	DELETE: Delete a post using DeletePost(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/posts", PostHandler)

In this example, PostHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/posts" on the server will be processed by PostHandler.
*/
func PostHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewPost(w, r)
		return
	case http.MethodGet:
		GetPosts(w, r)
		return
	case http.MethodPut:
		UpdatePost(w, r)
		return
	case http.MethodDelete:
		DeletePost(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
NewPost function is an HTTP handler for creating a new post in the web application.
This function extracts user data and post data from the HTTP request context, then inserts the new post into the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:

	Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
	Attempts to extract the post data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
	Calls controllers.InsertPost to insert the new post data into the database. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
	If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the post was successfully inserted.
*/
func NewPost(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newPostData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read post data from context", http.StatusInternalServerError)
		return
	}

	err := postcontrollers.InsertPost(dbutils.DB, userData.UserId, newPostData.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to insert post data", http.StatusInternalServerError)
		return
	}

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   newPostData,
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
GetPosts function is an HTTP handler for selecting all user viewable posts from the database.
This function extracts user data from the HTTP request context, then selects all posts a user should be able
to view from the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Calls controllers.SelectUserViewablePosts to select all viewable posts. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it writes the posts to a response and sends an HTTP 200 (OK) status to the client, indicating that the posts were successfully selected.
*/
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	specificUserId := r.URL.Query().Get("user_id")
	groupId := r.URL.Query().Get("group_id")

	var posts *dbmodels.Posts
	var err error

	if specificUserId != "" {
		posts, err = postcontrollers.SelectSpecificUserPosts(dbutils.DB, userInfo.UserId, specificUserId)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to select specific user posts", http.StatusInternalServerError)
			return
		}

	} else if groupId != "" {
		posts, err = postcontrollers.SelectGroupPosts(dbutils.DB, groupId)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to select user viewable posts", http.StatusInternalServerError)
			return
		}
	} else {
		posts, err = postcontrollers.SelectUserViewablePosts(dbutils.DB, userInfo.UserId)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to select user viewable posts", http.StatusInternalServerError)
			return
		}
	}

	sort.Slice(posts.Posts, func(i, j int) bool {
		return posts.Posts[i].PostInfo.CreationDate.After(posts.Posts[j].PostInfo.CreationDate)
	})

	response := readwritemodels.WriteData{
		Status: "success",
		Data:   posts,
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
UpdatePost function is an HTTP handler for updating a user's post in the database.
This function extracts user data and update post data from the HTTP request context, then updates a specific post
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the post data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.UpdateUserPost to update a post with new data. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the post was successfully updated.
*/
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	updatePostData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update post data from context", http.StatusInternalServerError)
		return
	}

	err := postcontrollers.UpdateUserPost(dbutils.DB, userData.UserId, updatePostData.Data)
	if err != nil {
		http.Error(w, "failed to update user post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
DeletePost function is an HTTP handler for deleting a user's post in the database.
This function extracts user data and delete post data from the HTTP request context, then delete a specific post
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the delete post data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.DeleteUserPost to delete a post. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the post was successfully deleted.
*/
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	deletePostData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read delete post data from context", http.StatusInternalServerError)
		return
	}

	err := postcontrollers.DeleteUserPost(dbutils.DB, userData.UserId, deletePostData.Data["post_id"].(string))
	if err != nil {
		http.Error(w, "failed to delete user post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
