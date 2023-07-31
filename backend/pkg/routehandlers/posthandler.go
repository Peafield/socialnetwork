package routehandlers

import (
	"net/http"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		NewPost(w, r)
		return
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	newPostData, ok := r.Context().Value(middleware.DataKey).(readwritemodels.ReadData)
	if !ok {
		http.Error(w, "failed to read post data from context", http.StatusInternalServerError)
		return
	}

	err := controllers.InsertPost(dbutils.DB, userInfo.UserId, newPostData.Data)
	if err != nil {
		http.Error(w, "failed to insert post data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserPosts(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(middleware.UserDataKey).(readwritemodels.Payload)
	if !ok {
		http.Error(w, "failed to read user data from context", http.StatusInternalServerError)
		return
	}

	userPosts, err := controllers.SelectPostsForUser(dbutils.DB, userInfo.UserId)
}
