package routehandlers

import (
	"net/http"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		//GetAllUsers(w, r)
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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	updateUserData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update post data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.UpdateUserAccount(dbutils.DB, userData.UserId, updateUserData.Data)
	if err != nil {
		http.Error(w, "failed to update user post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	// USERS SHOULD HAVE TO RETYPE THEIR CREDENTIALS TO DELETE THEIR OWN ACCOUNT
	deleteUserData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read update post data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.DeleteUserAccount(dbutils.DB, userData.UserId, deleteUserData.Data)
	if err != nil {
		http.Error(w, "failed to update user post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
