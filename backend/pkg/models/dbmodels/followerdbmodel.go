package dbmodels

import "time"

// Follower is a struct that holds follower data.
type Follower struct {
	FollowerId      string
	FollowingStatus int
	RequestPending  int
	CreationDate    time.Time
}

// Followers is a slice of Follower.
type Followers struct {
	Followers []Follower
}
