package dbmodels

import "time"

// User is a struct that holds user data.
type User struct {
	UserId         string    `json:"user_id"`
	IsLoggedIn     int       `json:"is_logged_in"`
	IsPrivate      int       `json:"is_private"`
	Email          string    `json:"email"`
	DisplayName    string    `json:"display_name"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DOB            time.Time `json:"dob"`
	AvatarPath     string    `json:"avatar_path"`
	AboutMe        string    `json:"about_me"`
	CreationDate   time.Time `json:"creation_date"`
}

type UserProfileData struct {
	UserInfo   User
	ProfilePic []byte
}

// Users is a slice of User
type Users struct {
	Users []User
}
