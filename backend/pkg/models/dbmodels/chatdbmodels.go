package dbmodels

import "time"

// Chat is a struct that holds chat data.
type Chat struct {
	ChatId       string    `json:"chat_id"`
	SenderId     string    `json:"sender_id"`
	ReceiverId   string    `json:"receiever_id"`
	GroupId      string    `json:"group_id"`
	CreationDate time.Time `json:"creation_date"`
}

// Chats is a slice of Chat.
type Chats struct {
	Chats []Chat
}

// ChatMessage is a struct that holds chat message data.
type ChatMessage struct {
	MessageId    string    `json:"message_id"`
	ChatId       string    `json:"chat_id"`
	SenderId     string    `json:"sender_id"`
	Message      string    `json:"message"`
	CreationDate time.Time `json:"creation_date"`
}

// ChatMessages is a slice of ChatMessage.
type ChatMessages struct {
	ChatMessages []ChatMessage
}
