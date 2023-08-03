package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"socialnetwork/pkg/controllers"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/models/dbmodels"
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
		postData["content"] = fmt.Sprintf("Content %v: Wow, I can't believe this content!", i)
		postData["privacy_level"] = rand.Intn(3)

		err := controllers.InsertPost(db, userIds[i-1], postData)
		if err != nil {
			return fmt.Errorf("failed to insert mock post data: %w", err)
		}
	}
	return nil
}
