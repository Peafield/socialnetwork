package routehandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialnetwork/pkg/controllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
SignInHandler is a HTTP handler function that processes sign in requests.

It expects the request context to contain a middleware.DataKey key that holds
a value of type readwritemodels.ReadData, which is extracted from the request
body by the ParseAndValidateData middleware.

The readwritemodels.ReadData struct should include a data field with the user
sign in details.

The handler does the following:

 1. Extracts form data from the context. If this operation fails, it returns an
    HTTP 500 (Internal Server Error) response.

 2. Calls controllers.ValidateCredentials function to validate the users credentials. If an error
    occurs, it returns an HTTP 401 Unauthorized response.

 3. Calls controllers.CreateWebToken function to create a JWT token for the newly
    signed in user. If the JWT creation fails, it returns an HTTP 500 response.

 4. If the previous steps succeed, the handler sends a success response containing
    the JWT token for the new user. This response is a JSON object of type
    readwritemodels.WriteData, with a "Status" field set to "success" and a "Data"
    field containing the JWT token.

HTTP Request Method: POST

Path: /signin

Example request body:

	{
		"status": "success",
		"data": {
			"username_email": "me@example.com",
			"password": "someHashedPassword",
		}
	}

Example response body on success:

	{
		"Status": "success",
		"Data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	}
*/
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//retrieve sign in form data
	formData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	//validate credentials
	user, err := controllers.ValidateCredentials(formData.Data, dbutils.DB)
	if err != nil {
		mi := fmt.Sprintf("failed to validate credentials, error: %s", err)
		http.Error(w, mi, http.StatusUnauthorized)
		return
	}

	//update user logged in status
	err = usercontrollers.UpdateLoggedInStatus(dbutils.DB, user.UserId, 1)
	if err != nil {
		http.Error(w, "failed to update logged in status", http.StatusInternalServerError)
		return
	}
	user.IsLoggedIn = 1

	//generate web token
	token, err := controllers.CreateWebToken(user)
	if err != nil {
		http.Error(w, "failed to create web token", http.StatusInternalServerError)
		return
	}

	//add token to response type, marshal and send back
	response := readwritemodels.WriteData{
		Status: "success",
		Data:   token,
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
