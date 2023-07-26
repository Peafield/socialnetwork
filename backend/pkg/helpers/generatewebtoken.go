package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"os"
	"socialnetwork/pkg/models/readwritemodels"
)

/*
GenerateWebToken creates a web token with a header, payload, and a signature.

First, we marshal the header and payload structs into bytes, then base64 encode them to strings.
We then join both of them, split by a full stop, called an unsignedToken.  A new hash is created using our secret key,
then the unsignedToken is written to this hash.  The bytes resulting from the sum of this hash is then base64 encoded
and appended to the unsignedToken after another full stop.

Parameters:
- header (readwritemodels.Header): a header type containing an algorithm type.
- payload (readwritemodels.Payload): a payload type containing information about the user and other information.

Returns:
- string: the generated web token.
- error

Errors:
- if there were errors marshalling the header or the payload.

Example:
- a web token is generated and assigned to a user when they login.
*/
func GenerateWebToken(header readwritemodels.Header, payload readwritemodels.Payload) (string, error) {
	headerstr, err := json.Marshal(header)
	if err != nil {
		return string(headerstr), err
	}
	header64 := base64.StdEncoding.EncodeToString(headerstr)

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		return string(payloadstr), err
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	unsignedToken := header64 + "." + payload64

	h := hmac.New(sha256.New, []byte(os.Getenv("SECRET_TOKEN_KEY")))

	h.Write([]byte(unsignedToken))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	token := unsignedToken + "." + signature

	return token, nil
}
