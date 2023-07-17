package dbmodels

import "time"

// User is a struct that holds user data.
type User struct {
	UserId         string
	IsLoggedIn     int
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	DOB            time.Time
	AvatarPath     string
	DisplayName    string
	AboutMe        string
	CreationDate   time.Time
}

// Users is a slice of User
type Users struct {
	Users []User
}
