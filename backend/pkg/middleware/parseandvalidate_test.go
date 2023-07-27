package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/readwritemodels"
	"strings"
	"testing"
)

const DataKey readwritemodels.ContextKey = iota

func TestParseAndValidateData(t *testing.T) {
	reqBody := `{"field1": "value1", "field2": "value2"}`
	req, err := http.NewRequest("POST", "/test", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Context().Value(DataKey)
		if value == nil {
			t.Error("Data not found in context")
		}
	})

	handler := middleware.ParseAndValidateData(nextHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	reqBody = `{"field1": "value1" "field2": "value2"}`
	req, err = http.NewRequest("POST", "/test", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
