package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	postcontrollers "socialnetwork/pkg/controllers/PostControllers"
	"strings"
)

func CreateFakePosts(db *sql.DB) error {
	postAmount := 20
	postData := map[string]interface{}{}

	userIds, err := GetAllUserIDs(db)
	if err != nil {
		return fmt.Errorf("failed to get all user id's when faking comments: %w", err)
	}

	for i := 1; i <= postAmount; i++ {
		content := ContentGenerator()
		postData["content"] = content
		postData["privacy_level"] = float64(rand.Intn(2))

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
