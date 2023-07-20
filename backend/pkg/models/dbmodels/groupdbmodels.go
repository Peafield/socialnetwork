package dbmodels

import "time"

// Group is a struct that holds group data.
type Group struct {
	GroupId      string
	Title        string
	Description  string
	CreatorId    string
	CreationDate time.Time
}

// Groups is a slice of Group.
type Groups struct {
	Groups []Group
}

// GroupMember is a struct that holds group member data.
type GroupMember struct {
	GroupId        string
	MemberId       string
	RequestPending int
	CreationDate   time.Time
}

// GroupMembers is a slice of GroupMember.
type GroupMembers struct {
	GroupMembers []GroupMember
}

// GroupEvent is a struct that holds group event data.
type GroupEvent struct {
	EventId        string
	GroupId        string
	CreatorId      string
	Title          string
	Description    string
	EventStartTime time.Time
	TotalGoing     int
	TotalNotGoing  int
	CreationDate   string
}

// GroupEvents is a slice of GroupEvent.
type GroupEvents struct {
	GroupEvents []GroupEvent
}

// GroupEventAttendee is a struct that holds group event attendee data.
type GroupEventAttendee struct {
	EventId         string
	AttendeeId      string
	AttendingStatus string
	EventStatus     int
	CreationDate    time.Time
}

// GroupEventAttendees is a slice of GroupEventAttendee.
type GroupEventAttendees struct {
	GroupEventAttendees []GroupEventAttendee
}
