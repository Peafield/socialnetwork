package controllers

import (
	"fmt"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"time"
)

func CreateWebToken(user *dbmodels.User) (string, error) {
	header := readwritemodels.Header{
		Alg: "sha256",
	}

	payload := readwritemodels.Payload{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      0,
		Exp:       time.Now().Add(time.Hour * 48).Unix(),
		Iat:       time.Now().Unix(),
	}
	token, err := helpers.GenerateWebToken(header, payload)
	if err != nil {
		return "", fmt.Errorf("failed to generate web token: %s", err)
	}

	return token, nil
}
