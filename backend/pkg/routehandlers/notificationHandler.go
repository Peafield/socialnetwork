package routehandlers

import (
	"log"
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
	log.Printf("hanlder: %s", newNotificationData)

	err := notificationcontrollers.InsertNewNotification(dbutils.DB, userData.UserId, newNotificationData.Data)
	if err != nil {
		http.Error(w, "failed to insert notification data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
