package controllers_test

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/controllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"testing"
	"time"
)

func TestValidateCredentials(t *testing.T) {
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
		RegisterData  map[string]interface{}
		SignInData    map[string]interface{}
		ExpectedError bool
	}{
		{
			Name: "Successfuly validated with email",
			RegisterData: map[string]interface{}{
				"email":        "test@",
				"password":     "PW",
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			SignInData: map[string]interface{}{
				"username_email": "test@",
				"password":       "PW",
			},
			ExpectedError: false,
		},
		{
			Name: "Successfuly validated with username",
			RegisterData: map[string]interface{}{
				"email":        "test@",
				"password":     "PW",
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			SignInData: map[string]interface{}{
				"username_email": "DN",
				"password":       "PW",
			},
			ExpectedError: false,
		},
		{
			Name: "Incorrect email",
			RegisterData: map[string]interface{}{
				"email":        "test@",
				"password":     "PW",
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			SignInData: map[string]interface{}{
				"username_email": "t3st@",
				"password":       "PW",
			},
			ExpectedError: true,
		},
		{
			Name: "Incorrect password",
			RegisterData: map[string]interface{}{
				"email":        "test@",
				"password":     "PW",
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			SignInData: map[string]interface{}{
				"username_email": "test@",
				"password":       "Pw",
			},
			ExpectedError: true,
		},
		{
			Name: "Incorrect username",
			RegisterData: map[string]interface{}{
				"email":        "test@",
				"password":     "PW",
				"first_name":   fmt.Sprintf("F%v", time.Now().Unix()),
				"last_name":    fmt.Sprintf("L%v", time.Now().Unix()),
				"dob":          time.Now(),
				"display_name": "DN",
				"about_me":     fmt.Sprintf("Test time: %v", time.Now().Unix()),
			},
			SignInData: map[string]interface{}{
				"username_email": "t3st",
				"password":       "PW",
			},
			ExpectedError: true,
		},
	}

	for i, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testingTime := time.Now().Add(time.Minute * time.Duration(i*10)).Unix()
			tc.RegisterData["email"] = fmt.Sprintf("test@%v", testingTime)
			tc.RegisterData["password"] = fmt.Sprintf("PW%v", testingTime)
			tc.RegisterData["display_name"] = fmt.Sprintf("DN%v", testingTime)

			tc.SignInData["username_email"] = fmt.Sprintf("%v%v", tc.SignInData["username_email"], testingTime)
			tc.SignInData["password"] = fmt.Sprintf("%v%v", tc.SignInData["password"], testingTime)

			_, err := usercontrollers.RegisterUser(db, tc.RegisterData)
			if err != nil {
				t.Logf("validating credentials testing register user error: %s", err)
			}

			_, err = controllers.ValidateCredentials(tc.SignInData, db)
			if tc.ExpectedError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.ExpectedError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
