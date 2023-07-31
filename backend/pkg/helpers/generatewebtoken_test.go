package helpers_test

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/readwritemodels"
	"strings"
	"testing"
	"time"
)

func TestGenerateWebToken(t *testing.T) {
	header := readwritemodels.Header{
		Alg: "sha256",
	}

	payload := readwritemodels.Payload{
		UserId:    "validUserId",
		FirstName: "Test",
		LastName:  "Tests",
		Role:      1,
		Exp:       time.Now().Unix(),
		Iat:       time.Now().Unix(),
	}

	validtoken, err := helpers.GenerateWebToken(header, payload)
	if err != nil {
		t.Errorf("failed to create webtoken: %s", err)
	}

	fullStopCount := strings.Count(validtoken, ".")
	if fullStopCount != 2 {
		t.Errorf("number of full stops = %v, but want 2", fullStopCount)
	}

	tokenHeader := strings.Split(validtoken, ".")[0]
	decodedTokenHeader, err := base64.StdEncoding.DecodeString(tokenHeader)
	if err != nil {
		t.Errorf("failed to decode token header: %s", err)
	}

	headerAsBytes, err := json.Marshal(header)
	if err != nil {
		t.Errorf("failed to marshal header: %s", err)
	}

	if !reflect.DeepEqual(decodedTokenHeader, headerAsBytes) {
		t.Errorf("headers are not equal to each other")
	}

	tokenPayload := strings.Split(validtoken, ".")[1]
	decodedTokenPayload, err := base64.StdEncoding.DecodeString(tokenPayload)
	if err != nil {
		t.Errorf("failed to decode token header: %s", err)
	}

	payloadAsBytes, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("failed to marshal payload: %s", err)
	}

	if !reflect.DeepEqual(decodedTokenPayload, payloadAsBytes) {
		t.Errorf("payloads are not equal to each other")
	}
}
