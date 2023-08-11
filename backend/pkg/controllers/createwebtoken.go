package controllers

import (
	"fmt"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"time"
)

/*
CreateWebToken is a function that generates a web token for a given user.

The function creates a Header struct and a Payload struct, both of which are used to generate a JWT (JSON Web Token).
The Payload struct is populated with relevant user information and some additional properties such as the token
expiration time and the time at which the token was issued.

Parameters:
  - user (*dbmodels.User): A pointer to the User struct from which user details are extracted for the payload.

Returns:
  - string: The generated web token.
  - error: An error is returned if there are any issues during the token generation process.

Operations:
  - Creates a Header struct with the algorithm set to "sha256"
  - Creates a Payload struct with various properties:
  - UserId: Extracted from the provided user
  - FirstName: Extracted from the provided user
  - LastName: Extracted from the provided user
  - Role: Set to 0 by default
  - Exp: Set to the current time plus 48 hours, represented as Unix timestamp
  - Iat: Set to the current time, represented as Unix timestamp
  - Calls GenerateWebToken with the created Header and Payload to generate the web token

Errors:
  - Returns an error if there's an issue generating the web token
*/
func CreateWebToken(user *dbmodels.User) (string, error) {
	header := readwritemodels.Header{
		Alg: "sha256",
	}

	payload := readwritemodels.Payload{
		UserId:      user.UserId,
		DisplayName: user.DisplayName,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		IsLoggedIn:  user.IsLoggedIn,
		Role:        0,
		Exp:         time.Now().Add(time.Hour * 48).Unix(),
		Iat:         time.Now().Unix(),
	}
	token, err := helpers.GenerateWebToken(header, payload)
	if err != nil {
		return "", fmt.Errorf("failed to generate web token: %s", err)
	}

	return token, nil
}
