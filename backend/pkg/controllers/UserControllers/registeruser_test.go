package usercontrollers_test

import (
	"database/sql"
	"fmt"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"testing"
	"time"
)

func TestRegisterUser(t *testing.T) {
	db, err := sql.Open("sqlite3", "../db/testDB.db")
	if err != nil {
		t.Errorf("failed to open database: %s", err)
	} else {
		defer db.Close()
	}

	err = dbstatements.InitDBStatements(db)
	if err != nil {
		t.Errorf("Failed to prepare database statements: %s", err)
	} else {
		defer dbstatements.CloseDBStatements()
	}

	testCases := []struct {
		Name          string
		Status        string
		Data          map[string]interface{}
		ExpectedError bool
	}{
		{
			Name:   "Successfuly insertion",
			Status: "success",
			Data: map[string]interface{}{
				"email":        "test@",
				"password":     fmt.Sprintf("PW%v", time.Now().Unix()),
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			ExpectedError: false,
		},
		{
			Name:   "Successfuly insertion",
			Status: "success",
			Data: map[string]interface{}{
				"email":        "test@",
				"password":     fmt.Sprintf("PW%v", time.Now().Unix()),
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			ExpectedError: false,
		},
		{
			Name:   "Bad data",
			Status: "success",
			Data: map[string]interface{}{
				"email":        1,
				"password":     fmt.Sprintf("PW%v", time.Now().Unix()),
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    time.Now().Unix(),
				"dob":          "DN",
				"display_name": fmt.Sprintf("DN%v", time.Now().Unix()),
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			ExpectedError: true,
		},
		{
			Name:   "Non-unqiue email",
			Status: "success",
			Data: map[string]interface{}{
				"email":        "user@test.com",
				"password":     fmt.Sprintf("PW%v", time.Now().Unix()),
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			ExpectedError: true,
		},
		{
			Name:   "Non-unqiue display name",
			Status: "success",
			Data: map[string]interface{}{
				"email":        "test@",
				"password":     fmt.Sprintf("PW%v", time.Now().Unix()),
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "user6",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			ExpectedError: true,
		},
		{
			Name:          "Missing user data",
			Status:        "success",
			Data:          map[string]interface{}{},
			ExpectedError: true,
		},
	}

	for i, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testingTime := time.Now().Add(time.Minute * time.Duration(i*10)).Unix()
			if tc.Name != "Non-unqiue email" && tc.Name != "Non-unqiue display name" {
				tc.Data["email"] = fmt.Sprintf("test@%v", testingTime)
				tc.Data["password"] = fmt.Sprintf("PW%v", testingTime)
				tc.Data["display_name"] = fmt.Sprintf("DN%v", testingTime)
			}
			_, err := usercontrollers.RegisterUser(tc.Data, db, dbstatements.InsertUserStmt)
			if tc.ExpectedError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.ExpectedError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
