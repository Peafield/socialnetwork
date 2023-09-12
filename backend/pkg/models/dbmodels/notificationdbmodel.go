package dbmodels

import "time"

// Notification is a struct that holds notification data.
type Notification struct {
	NotificationId string    `json:"notification_id"`
	SenderId       string    `json:"sender_id"`
	ReceiverId     string    `json:"receiver_id"`
	GroupId        string    `json:"group_id"`
	PostId         string    `json:"post_id"`
	EventId        string    `json:"event_id"`
	CommentId      string    `json:"comment_id"`
	ChatId         string    `json:"chat_id"`
	ActionType     string    `json:"action_type"`
	ReadStatus     int       `json:"read_status"`
	CreationDate   time.Time `json:"creation_date"`
}

// Notifications is a slice of Notification
type Notifications struct {
	Notifications []Notification
}
