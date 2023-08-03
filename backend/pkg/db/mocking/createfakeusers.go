package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbstatements"
	"strings"
	"time"
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
		formData["email"] = fakeNameData.EmailU + "@" + fakeNameData.EmailD
		formData["password"] = "abc123"

		nameSplit := strings.Split(fakeNameData.Name, " ")

		formData["first_name"] = nameSplit[0]
		formData["last_name"] = nameSplit[1]
		formData["display_name"] = fakeNameData.UserName

		dob, err := time.Parse("2006-01-02", fakeNameData.BirthData)
		if err != nil {
			log.Fatal(err)
		}

		formData["dob"] = dob
		formData["about_me"] = fmt.Sprintf("Hi my name is %s! I'm excited to meet you", fakeNameData.Name)

		_, err = controllers.RegisterUser(formData, db, dbstatements.InsertUserStmt)
		if err != nil {
			return fmt.Errorf("failed to insert fake users: %w", err)
		}
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
