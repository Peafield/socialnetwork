package dbmodels

import "time"

// Reaction is a struct that holds reaction data.
type Reaction struct {
	UserId       string
	PostId       string
	CommentId    string
	Reaction     int
	CreationDate time.Time
}

// Reactions is a slice of reaction
type Reactions struct {
	Reactions []Reaction
}
