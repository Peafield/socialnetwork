package routehandlers

import (
	"encoding/json"
	"net/http"
	notificationcontrollers "socialnetwork/pkg/controllers/NotificationControllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewNotification(w, r)
		return
	case http.MethodDelete:
		DeleteNotification(w, r)
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

}

func NewNotification(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newNotificationData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read notification data from context", http.StatusInternalServerError)
		return
	}

	err := notificationcontrollers.InsertNewNotification(dbutils.DB, userData.UserId, newNotificationData.Data)
	if err != nil {
		http.Error(w, "failed to insert notification data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	deleteNotificationData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read notification data from context", http.StatusInternalServerError)
		return
	}

	err := notificationcontrollers.DeleteNotification(dbutils.DB, userData.UserId, deleteNotificationData.Data)
	if err != nil {
		http.Error(w, "failed to delete user post", http.StatusInternalServerError)
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
