package routehandlers

import (
	"encoding/json"
	"net/http"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	formData, ok := r.Context().Value(middleware.DataKey).(*readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read form data from context", http.StatusInternalServerError)
		return
	}

	userId, err := controllers.RegisterUser(formData)
	if err != nil {
		http.Error(w, "failed to reigster user", http.StatusInternalServerError)
		return
	}

	token, err := controllers.CreateWebToken(userId)
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
