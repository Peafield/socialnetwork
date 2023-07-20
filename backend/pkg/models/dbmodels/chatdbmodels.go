package dbmodels

import "time"

// Chat is a struct that holds chat data.
type Chat struct {
	ChatId       string
	SenderId     string
	ReceiverId   string
	CreationDate time.Time
}

// Chats is a slice of Chat.
type Chats struct {
	Chats []Chat
}

// ChatMessage is a struct that holds chat message data.
type ChatMessage struct {
	MessageId    string
	ChatId       string
	SenderId     string
	Message      string
	Timestamp    time.Time
	CreationDate time.Time
}

// ChatMessages is a slice of ChatMessage.
type ChatMessages struct {
	ChatMessages []ChatMessage
}
