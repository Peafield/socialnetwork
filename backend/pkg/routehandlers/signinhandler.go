package routehandlers

import (
	"encoding/json"
	"net/http"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
 */
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve sign in form data
	formData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	//validate credentials
	user, err := controllers.ValidateCredentials(formData.Data)
	if err != nil {
		http.Error(w, "failed to validate credentials", http.StatusUnauthorized)
		return
	}

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
