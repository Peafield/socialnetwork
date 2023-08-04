package dbmodels

import "time"

// Follower is a struct that holds follower data.
type PostSelectedFollower struct {
	PostId            string    `json:"post_id"`
	AllowedFollowerId string    `json:"allowed_follower_id"`
	CreationDate      time.Time `json:"creation_date"`
}

// Followers is a slice of Follower.
type PostSelectedFollowers struct {
	PostSelectedFollowers []PostSelectedFollower
}
