package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"strings"
)

const (
	UserDataKey readwritemodels.ContextKey = iota
	DataKey
)

/*
ValidateTokenMiddleware is a middleware function that validates a webtoken included in the Authorization header of a http request.

If the header is valid and contains the bearer token, the bearer token is then verified. If the bearer token is valid, the payload is then
written into the request context to be passed on in the chain to either parse and validate the data or continue to a route handler.

Parameters:
  - next http.Handler: The next handler in the middleware chain to which the request should be passed if the token is valid.

Return:
  - http.HandlerFunc: An HTTP handler that wraps the input handler (next) with token validation logic.

Example:
  - When the client sends a post request, the user's authorisation details will also be passed in the header to be validated.
*/
func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		authType, token, found := strings.Cut(authorizationHeader, " ")
		if !found || authType != "Bearer" {
			http.Error(w, "Invalid or missing authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := token

		validToken, err := VerifyToken(bearerToken)
		if !validToken || err != nil {
			http.Error(w, "Invalid token verification", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(bearerToken, ".")
		payloadEncoded := tokenParts[1]
		payload, err := base64.StdEncoding.DecodeString(payloadEncoded)
		if err != nil {
			http.Error(w, "failed to decode token", http.StatusUnauthorized)
			return
		}

		var recievedPayload readwritemodels.Payload
		err = json.Unmarshal(payload, &recievedPayload)
		if err != nil {
			http.Error(w, "Invalid token when unmarshalling into payload struct", http.StatusUnauthorized)
			return
		}

		err = ValidateLoggedInStatus(recievedPayload)
		if err != nil {
			r.Header.Del("Authorization")
			http.Error(w, error.Error(err), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserDataKey, recievedPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Validating "middleware" for websocket
func ValidateTokenWebSocket(authorizationHeader string) (*readwritemodels.Payload, error) {
	authType, token, found := strings.Cut(authorizationHeader, " ")
	if !found || authType != "Bearer" {
		return nil, fmt.Errorf("invalid or missing authorization header")
	}

	bearerToken := token

	validToken, err := VerifyToken(bearerToken)
	if !validToken || err != nil {
		return nil, fmt.Errorf("Invalid token verification")
	}

	tokenParts := strings.Split(bearerToken, ".")
	payloadEncoded := tokenParts[1]
	payload, err := base64.StdEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token")
	}

	var recievedPayload readwritemodels.Payload
	err = json.Unmarshal(payload, &recievedPayload)
	if err != nil {
		return nil, fmt.Errorf("Invalid token when unmarshalling into payload struct")
	}

	err = ValidateLoggedInStatus(recievedPayload)
	if err != nil {
		return nil, fmt.Errorf(error.Error(err))
	}
	return &recievedPayload, nil
}

/*
VerifyToken confirms with a token is valid or not.

The function takes a token which is then split and decoded into three parts.
The header and the payload or then re-endecoded with a hashed version of secret token key
to check whether this new token and the inputted token are the same. If they are true, is returned, else
the token is invalid and false is returned.

Parameters:
  - token (string): a token.

Returns:
  - bool: True if the token is valid, false if not.
  - error: An error will be returned if the token does not have three parts or if the header or payload cannot be decoded.
*/
func VerifyToken(token string) (bool, error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return false, nil
	}

	unsignedStr := splitToken[0] + "." + splitToken[1]
	h := hmac.New(sha256.New, []byte(os.Getenv("SECRET_TOKEN_KEY")))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != splitToken[2] {
		return false, nil
	}

	return true, nil
}

func ValidateLoggedInStatus(payload readwritemodels.Payload) error {
	queryValues := []interface{}{
		payload.UserId,
	}

	userData, err := crud.SelectFromDatabase(dbutils.DB, "Users", dbstatements.SelectUserByIDStmt, queryValues)
	if err != nil {
		return fmt.Errorf("failed to select user in validate logged in status: %w", err)
	}

	user, ok := userData[0].(*dbmodels.User)
	if !ok {
		return fmt.Errorf("failed to assert user type when validating logged in status")
	}

	if user.IsLoggedIn != payload.IsLoggedIn {
		return fmt.Errorf("mismatch in logged in status")
	}
	return nil
}
