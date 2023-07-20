package dbmodels

import "time"

// Post is a struct that holds post data.
type Post struct {
	PostId       string
	GroupId      string
	CreatorId    string
	Title        string
	ImagePath    string
	Content      string
	PrivacyLevel int
	// Allowed Followers should be a slice of followers
	AllowedFollowers string
	Likes            int
	Dislikes         int
	CreationDate     time.Time
}

// Posts is slice of Post.
type Posts struct {
	Posts []Post
}
