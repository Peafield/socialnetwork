package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"socialnetwork/pkg/models/readwritemodels"
	"strings"
)

type key int

const PayloadKey key = iota

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

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Invalid or missing authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := headerParts[1]

		validToken, err := VerifyToken(bearerToken)
		if !validToken || err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(bearerToken, ".")
		payloadEncoded := tokenParts[1]
		payload, err := base64.StdEncoding.DecodeString(payloadEncoded)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		var RecievedPayload readwritemodels.Payload
		err = json.Unmarshal(payload, &RecievedPayload)
		if err != nil {
			http.Error(w, "Invalid token when unmarshalling into payload struct", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), PayloadKey, RecievedPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(os.Getenv("SECRET_TOKEN_KEY")))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != splitToken[2] {
		return false, nil
	}

	return true, nil
}
