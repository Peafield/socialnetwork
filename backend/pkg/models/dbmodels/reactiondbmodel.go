package dbmodels

import "time"

// Reaction is a struct that holds reaction data.
type Reaction struct {
	UserId       string    `json:"user_id"`
	PostId       *string   `json:"post_id"`
	CommentId    *string   `json:"comment_id"`
	Reaction     string    `json:"reaction"`
	CreationDate time.Time `json:"creation_date"`
}

// Reactions is a slice of reaction
type Reactions struct {
	Reactions []Reaction
}
