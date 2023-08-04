package dbmodels

import "time"

// Follower is a struct that holds follower data.
type Follower struct {
	FollowerId      string    `json:"follower_id"`
	FolloweeId      string    `json:"followee_id"`
	FollowingStatus int       `json:"following_status"`
	RequestPending  int       `json:"request_pending"`
	CreationDate    time.Time `json:"creation_date"`
}

// Followers is a slice of Follower.
type Followers struct {
	Followers []Follower
}
