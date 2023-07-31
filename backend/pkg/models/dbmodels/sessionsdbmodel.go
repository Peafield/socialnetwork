package dbmodels

import "time"

// Session is a struct that holds session data.
type Session struct {
	SessionId    string    `json:"session"`
	UserId       string    `json:"user_id"`
	CreationDate time.Time `json:"creation_date"`
}
