package dbmodels

import "time"

// User is a struct that holds user data.
type User struct {
	UserId         string    `json:"user_id"`
	IsLoggedIn     int       `json:"is_logged_in"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DOB            time.Time `json:"dob"`
	AvatarPath     string    `json:"avatar_path"`
	DisplayName    string    `json:"display_name"`
	AboutMe        string    `json:"about_me"`
	CreationDate   time.Time `json:"creation_date"`
}

// Users is a slice of User
type Users struct {
	Users []User
}
