package controllers_test

import (
	"fmt"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/models/dbmodels"
	"testing"
	"time"
)

func TestCreateWebToken(t *testing.T) {
	user := &dbmodels.User{
		UserId:    fmt.Sprintf("userId%v", time.Now().Unix()),
		FirstName: fmt.Sprintf("FName%v", time.Now().Unix()),
		LastName:  fmt.Sprintf("LName%v", time.Now().Unix()),
	}
	token, err := controllers.CreateWebToken(user)
	if err != nil {
		t.Errorf("error creating web token: %s", err)
	}

	if len(token) == 0 {
		t.Errorf("token failed to be created")
	}
}
