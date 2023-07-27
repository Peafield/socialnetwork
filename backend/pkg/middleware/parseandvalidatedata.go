package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"socialnetwork/pkg/models/readwritemodels"
)

const DataKey readwritemodels.ContextKey = iota

/*
ParseAndValidateData is a middleware function that extracts and decodes JSON data from the request body.

It then stores the decoded data into the request's context for further handlers to use. If the data cannot be
parsed, it sends an HTTP error response with a status code of 400 (Bad Request).

Parameters:
  - next (http.Handler): The next handler to be called in the middleware chain.

Returns:
  - http.Handler: An HTTP handler that processes the incoming request.

Errors:
  - If the request payload is invalid, sends an HTTP error with status code 400 (Bad Request).
*/
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
