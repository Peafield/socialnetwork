package routehandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"socialnetwork/pkg/controllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

var ErrUserExists = errors.New("user display name or email already in use")

/*
SignUpHandler is a HTTP handler function that processes sign up requests.

It expects the request context to contain a middleware.DataKey key that holds
a value of type readwritemodels.ReadData, which is extracted from the request
body by the ParseAndValidateData middleware.

The readwritemodels.ReadData struct should include a data field with the user
registration details.

The handler does the following:

 1. Extracts form data from the context. If this operation fails, it returns an
    HTTP 500 (Internal Server Error) response.

 2. Calls controllers.RegisterUser function to register the new user. If an error
    occurs (e.g., because of a database issue or an issue with the provided data),
    it returns an HTTP 500 response.

 3. Calls controllers.CreateWebToken function to create a JWT token for the newly
    registered user. If the JWT creation fails, it returns an HTTP 500 response.

 4. If the previous steps succeed, the handler sends a success response containing
    the JWT token for the new user. This response is a JSON object of type
    readwritemodels.WriteData, with a "Status" field set to "success" and a "Data"
    field containing the JWT token.

HTTP Request Method: POST

Path: /signup

Example request body:

	{
		"status": "success",
		"data": {
			"email": "me@example.com",
			"password": "someHashedPassword",
			"first_name": "John",
			"last_name": "Doe",
			"dob": "2000-01-01T00:00:00Z",
			"display_name": "johnny",
			"about_me": "Hello, I am John Doe!"
		}
	}

Example response body on success:

	{
		"Status": "success",
		"Data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	}
*/
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	formData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	/*IMAGE HANDLING CODE*/ //////////////////////////////////////////////////////

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // Limit the size of uploaded files
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, handler, err := r.FormFile("avatar_path")
	if err != nil {
		http.Error(w, "Unable to get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server to store the uploaded image
	filepath := "backend/pkg/db/images/" + handler.Filename
	newFile, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file's content to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file content", http.StatusInternalServerError)
		return
	}

	/////////////////////////////////////////////////////////////////////

	user, err := usercontrollers.RegisterUser(formData.Data, dbutils.DB, dbstatements.InsertUserStmt)
	if err != nil {
		if errors.Is(err, ErrUserExists) {
			http.Error(w, "user display name or email already in use", http.StatusBadRequest)
		} else {
			fmt.Println(err)
			http.Error(w, "failed to reigster user", http.StatusInternalServerError)
		}
		return
	}

	token, err := controllers.CreateWebToken(user)
	if err != nil {
		http.Error(w, "failed to create web token", http.StatusInternalServerError)
		return
	}

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
