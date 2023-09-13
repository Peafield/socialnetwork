package routehandlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	reactioncontrollers "socialnetwork/pkg/controllers/ReactionControllers"
	"socialnetwork/pkg/db/dbutils"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
ReactionHandler is a general HTTP request handler for actions related to reactions.
It is designed to process different HTTP methods (GET, POST, PUT, DELETE) and call the corresponding functions for each method.

Based on the HTTP method, it will:

	POST: Create a new reaction using NewReaction(w, r)
	GET: Retrieve a posts reactions using PostCommentReactions(w, r)
	PUT: Update an existing reaction using UpdateReaction(w, r)

If the request's HTTP method is not one of the above, the function will respond with an HTTP 400 (Bad Request) status,
indicating that the server cannot or will not process the request due to something perceived to be a client error.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Usage:

http.HandleFunc("/reactions", ReactionHandler)

In this example, ReactionHandler is registered with the HTTP package's default ServeMux (which is a HTTP request router).
This means that any HTTP request sent to the path "/reactions" on the server will be processed by ReactionHandler.
*/
func ReactionHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		HandleReaction(w, r)
		return
	case http.MethodGet:
		GetReaction(w, r)
		return
	// case http.MethodPut:
	// 	UpdateReaction(w, r)
	// 	return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

/*
NewReaction is an HTTP handler for creating a new reaction in the web application.
This function extracts user data and reaction data from the HTTP request context, then inserts the new reaction into the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the reaction data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the reaction data from the context.
  - Calls controllers.InsertReaction to insert the new reaction data into the database. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the reaction data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the reaction was successfully inserted.
*/
func HandleReaction(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newReactionData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read reaction data from context", http.StatusInternalServerError)
		return
	}

	err := reactioncontrollers.HandleUserReaction(dbutils.DB, userData.UserId, newReactionData.Data)
	if err != nil {
		http.Error(w, "failed to insert reaction data", http.StatusInternalServerError)
		return
	}

	response := readwritemodels.WriteData{
		Status: "success",
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

func GetReaction(w http.ResponseWriter, r *http.Request) {
	var response readwritemodels.WriteData

	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	postId := r.URL.Query().Get("post_id")
	commentId := r.URL.Query().Get("comment_id")

	if postId != "" {
		userReaction, err := reactioncontrollers.GetUserPostReaction(dbutils.DB, userData.UserId, postId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			log.Println(err)
			http.Error(w, "failed to get user reaction", http.StatusInternalServerError)
			return
		}

		response = readwritemodels.WriteData{
			Status: "success",
			Data:   userReaction,
		}
	} else if commentId != "" {
		userReaction, err := reactioncontrollers.GetUserCommentReaction(dbutils.DB, userData.UserId, commentId)
		if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
			log.Println(err)
			http.Error(w, "failed to get user reaction", http.StatusInternalServerError)
			return
		}

		response = readwritemodels.WriteData{
			Status: "success",
			Data:   userReaction,
		}
	} else {
		response = readwritemodels.WriteData{
			Status: "success",
		}
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
PostCommentReactions is an HTTP handler for selecting all reactions or comments for a post or a comment from the database.
This function extracts reaction data from the HTTP request context, then selects all related reactions rom the database

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the reaction data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Calls controllers.SelectReactions to select all related reactions. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the reaction data.
  - If all of the above steps are successful, it writes the reactions to a response and sends an HTTP 200 (OK) status to the client, indicating that the reactions were successfully selected.
*/
// func PostCommentReactions(w http.ResponseWriter, r *http.Request) {
// 	reactionData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read reaction id data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	reactions, err := reactioncontrollers.SelectReactions(dbutils.DB, reactionData.Data)
// 	if err != nil {
// 		http.Error(w, "failed to select reactions", http.StatusInternalServerError)
// 		return
// 	}

// 	response := readwritemodels.WriteData{
// 		Status: "success",
// 		Data:   reactions,
// 	}
// 	jsonReponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonReponse)

// }

/*
UpdateReaction is an HTTP handler for updating a reaction in the database.
This function extracts user data and update reaction data from the HTTP request context, then updates a specific reaction
based on that data in the database.

Parameters:
  - w (http.ResponseWriter): An HTTP ResponseWriter interface that forms the response that will be written to the HTTP connection.
  - r (*http.Request): A pointer to the HTTP request received from the client.

Specifically, it:
  - Attempts to extract the user data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the user data from the context.
  - Attempts to extract the update reaction data from the context of the request. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to read the post data from the context.
  - Calls controllers.UpdateUserReaction to update a reaction. If it fails, it sends an HTTP 500 (Internal Server Error) status to the client, with an error message indicating that it failed to insert the post data.
  - If all of the above steps are successful, it sends an HTTP 200 (OK) status to the client, indicating that the reaction was successfully updated.
*/
// func UpdateReaction(w http.ResponseWriter, r *http.Request) {
// 	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
// 	if !ok {
// 		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	updateReactionData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
// 	if !ok {
// 		http.Error(w, "failed to read update reaction data from context", http.StatusInternalServerError)
// 		return
// 	}

// 	err := reactioncontrollers.UpdateUserPostOrCommentReaction(dbutils.DB, userData.UserId, updateReactionData.Data)
// 	if err != nil {
// 		http.Error(w, "failed to update user reaction", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
