package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
	"testing"
	"time"
)

func TestValidateTokenMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test passed")) // If we reach this point, the test passed
	})

	handler := middleware.ValidateTokenMiddleware(nextHandler)

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "No Authorization header",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Authorization Format",
			token:          "Bearer token gibberish",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Authorization scheme",
			token:          "Basic invalid.token",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Authorization header",
			token:          "Bearer invalid.token",
			expectedStatus: http.StatusUnauthorized,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			if tt.token != "" {
				request.Header.Set("Authorization", tt.token)
			}

			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, request)

			if status := recorder.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
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

	testCases := []struct {
		Name          string
		InputToken    string
		ValidateToken string
		ExpectedError bool
	}{
		{
			Name:          "Valid token",
			InputToken:    validtoken,
			ValidateToken: validtoken,
			ExpectedError: false,
		},
		{
			Name:          "Invalid input token",
			InputToken:    "Invalid input token",
			ValidateToken: validtoken,
			ExpectedError: true,
		},
		{
			Name:          "Invalid Validation token",
			InputToken:    validtoken,
			ValidateToken: "invalid validation token",
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			IsValidToken, err := middleware.VerifyToken(tc.InputToken)
			if tc.ExpectedError && IsValidToken {
				t.Error("Expected invalid token, but got an valid one")
			} else if !tc.ExpectedError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
