package dbmodels

import "time"

// Notification is a struct that holds notification data.
type Notification struct {
	NotificationId string
	SenderId       string
	ReceiverId     string
	GroupId        string
	PostId         string
	EventId        string
	CommentId      string
	ChatId         string
	ReactionType   string
	ReadStatus     int
	CreationDate   time.Time
}

// Notifications is a slice of Notification
type Notifications struct {
	Notifications []Notification
}
