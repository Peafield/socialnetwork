package dbmodels

import "time"

// Post is a struct that holds post data.
type Post struct {
	PostId        string    `json:"post_id"`
	GroupId       string    `json:"group_id"`
	CreatorId     string    `json:"creator_id"`
	ImagePath     string    `json:"image_path"`
	Content       string    `json:"content"`
	NumOfComments int       `json:"num_of_comments"`
	PrivacyLevel  int       `json:"privacy_level"`
	Likes         int       `json:"likes"`
	Dislikes      int       `json:"dislikes"`
	CreationDate  time.Time `json:"creation_date"`
}

type PostData struct {
	PostInfo    Post
	PostPicture []byte
}

// Posts is slice of Post.
type Posts struct {
	Posts []PostData
}
