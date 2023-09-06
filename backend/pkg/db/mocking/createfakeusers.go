package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/pkg/controllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"strings"
)

type FakeName struct {
	Name      string `json:"name"`
	EmailU    string `json:"email_u"`
	EmailD    string `json:"email_d"`
	UserName  string `json:"username"`
	BirthData string `json:"birth_data"`
}

func CreateFakeUsers(db *sql.DB) error {
	userAmount := 20
	formData := map[string]interface{}{}
	for i := 1; i <= userAmount; i++ {
		fakeNameData := GetFakeUser()
		nameSplit := strings.Split(fakeNameData.Name, " ")
		nameSplit[0] = strings.ReplaceAll(nameSplit[0], ".", "")
		formData["first_name"] = nameSplit[0]
		formData["last_name"] = nameSplit[1]
		formData["email"] = nameSplit[0] + "." + nameSplit[1] + "@" + fakeNameData.EmailD
		formData["password"] = "abc123"

		formData["display_name"] = fakeNameData.UserName

		formData["dob"] = fakeNameData.BirthData
		formData["about_me"] = fmt.Sprintf("Hi my name is %s! I'm excited to meet you", fakeNameData.Name)

		_, err := usercontrollers.RegisterUser(db, formData)
		if err != nil {
			return fmt.Errorf("failed to insert fake users: %w", err)
		}
	}
	err := controllers.SignOutAllUsers(db)
	if err != nil {
		return fmt.Errorf("failed to sign out all users during mocking: %w", err)
	}

	return nil
}

func GetFakeUser() FakeName {
	fakeNameGen := "https://api.namefake.com/"
	response, err := http.Get(fakeNameGen)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var fakeName FakeName
	err = json.NewDecoder(response.Body).Decode(&fakeName)
	if err != nil {
		log.Fatal(err)
	}
	return fakeName
}
