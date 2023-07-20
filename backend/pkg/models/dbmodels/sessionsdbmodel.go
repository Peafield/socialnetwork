package dbmodels

import "time"

// Session is a struct that holds session data.
type Session struct {
	SessionId    string
	UserId       string
	CreationDate time.Time
}
