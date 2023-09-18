package dbmodels

import "time"

// Comment is a struct that holds comment data.
type Comment struct {
	CommentId    string    `json:"comment_id"`
	UserId       string    `json:"user_id"`
	PostId       string    `json:"post_id"`
	Content      string    `json:"content"`
	ImagePath    string    `json:"image_path"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	CreationDate time.Time `json:"creation_date"`
}

type CommentData struct {
	CommentInfo    Comment
	CommentPicture []byte
}

// Comments is a slice of Comment.
type Comments struct {
	Comments []CommentData
}
