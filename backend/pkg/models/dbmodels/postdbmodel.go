package dbmodels

import "time"

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

type Posts struct {
	Posts []Post
}
