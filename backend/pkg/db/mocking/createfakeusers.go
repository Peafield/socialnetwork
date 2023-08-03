package db

import (
	"database/sql"
	"fmt"
	"socialnetwork/pkg/controllers"
	"socialnetwork/pkg/db/dbstatements"
	"time"
)

func CreateFakeUsers(db *sql.DB) error {
	userAmount := 20
	formData := map[string]interface{}{}
	for i := 1; i <= userAmount; i++ {
		formData["email"] = fmt.Sprintf("testUser%v@mail.com", i)
		formData["password"] = fmt.Sprintf("password%v", i)
		formData["first_name"] = fmt.Sprintf("FNtestUser%v", i)
		formData["last_name"] = fmt.Sprintf("LNtestUser%v", i)
		formData["display_name"] = fmt.Sprintf("testUser%v", i)
		formData["dob"] = time.Now()
		formData["about_me"] = fmt.Sprintf("Hi my name is %s, I'm excited to meet you", formData["display_name"])

		_, err := controllers.RegisterUser(formData, db, dbstatements.InsertUserStmt)
		if err != nil {
			return fmt.Errorf("failed to insert fake users: %w", err)
		}
	}

	return nil
}
