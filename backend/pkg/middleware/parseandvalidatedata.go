package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"socialnetwork/pkg/models/readwritemodels"
)

const DataKey readwritemodels.ContextKey = iota

func ParseAndValidateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data readwritemodels.ReadData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), DataKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
