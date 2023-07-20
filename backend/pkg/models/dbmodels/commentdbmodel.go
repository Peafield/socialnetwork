package dbmodels

import "time"

// Comment is a struct that holds comment data.
type Comment struct {
	CommentId    string
	UserId       string
	PostId       string
	Content      string
	ImagePath    string
	Likes        int
	Dislikes     int
	Timestamp    time.Time
	CreationDate time.Time
}

// Comments is a slice of Comment.
type Comments struct {
	Comments []Comment
}
