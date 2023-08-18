package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	postcontrollers "socialnetwork/pkg/controllers/PostControllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
	"strings"
)

func CreateFakePosts(db *sql.DB) error {
	postAmount := 20
	postData := map[string]interface{}{}
	queryStatement := `
	SELECT * FROM Users
	`
	userIDData, err := crud.SelectFromDatabase(db, "Users", queryStatement, []interface{}{})
	if err != nil {
		return fmt.Errorf("failed to get userId's from db: %w", err)
	}

	var userIds []string
	for _, v := range userIDData {
		if user, ok := v.(*dbmodels.User); ok {
			userIds = append(userIds, user.UserId)
		} else {
			return fmt.Errorf("failed to assert user id data")
		}
	}

	for i := 1; i <= postAmount; i++ {
		content := ContentGenerator()
		postData["content"] = content
		postData["privacy_level"] = rand.Intn(3)

		err := postcontrollers.InsertPost(db, userIds[i-1], postData)
		if err != nil {
			return fmt.Errorf("failed to insert mock post data: %w", err)
		}
	}
	return nil
}

func ContentGenerator() string {
	loremIpsumGen := "https://loripsum.net/api/2/short/plaintext"
	response, err := http.Get(loremIpsumGen)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(body)
	pos := strings.Index(s, ".")

	return s[pos+1:]
}
