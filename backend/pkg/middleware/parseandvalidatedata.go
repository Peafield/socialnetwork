package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"socialnetwork/pkg/models/readwritemodels"
)

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
// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/signup")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
// }

func ParseAndValidateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data readwritemodels.ReadData
		err := json.NewDecoder(r.Body).Decode(&data)
		if r.Method != http.MethodGet && err != nil {
			fmt.Println(err)
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), DataKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
