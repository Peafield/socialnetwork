package dbmodels

import "time"

// Post is a struct that holds post data.
type Post struct {
	PostId       string `json:"post_id"`
	GroupId      string `json:"group_id"`
	CreatorId    string `json:"creator_id"`
	Title        string `json:"title"`
	ImagePath    string `json:"image_path"`
	Content      string `json:"content"`
	PrivacyLevel int    `json:"privacy_level"`
	// Allowed Followers should be a slice of followers
	AllowedFollowers string    `json:"allowed_followers"`
	Likes            int       `json:"likes"`
	Dislikes         int       `json:"dislikes"`
	CreationDate     time.Time `json:"creation_date"`
}

// Posts is slice of Post.
type Posts struct {
	Posts []Post
}
