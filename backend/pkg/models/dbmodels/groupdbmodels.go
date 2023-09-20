package dbmodels

import "time"

// Group is a struct that holds group data.
type Group struct {
	GroupId      string    `json:"group_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatorId    string    `json:"creator_id"`
	CreationDate time.Time `json:"creation_date"`
}

// Groups is a slice of Group.
type Groups struct {
	Groups []Group
}

// GroupMember is a struct that holds group member data.
type GroupMember struct {
	GroupId         string    `json:"group_id"`
	MemberId        string    `json:"member_id"`
	RequestPending  int       `json:"request_pending"`
	PermissionLevel int       `json:"permission_level"`
	CreationDate    time.Time `json:"creation_date"`
}

// GroupMembers is a slice of GroupMember.
type GroupMembers struct {
	GroupMembers []GroupMember
}

// GroupEvent is a struct that holds group event data.
type GroupEvent struct {
	EventId        string    `json:"event_id"`
	GroupId        string    `json:"group_id"`
	CreatorId      string    `json:"creator_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	EventStartTime time.Time `json:"event_start_time"`
	TotalGoing     int       `json:"total_going"`
	TotalNotGoing  int       `json:"total_not_going"`
	CreationDate   string    `json:"creation_date"`
}

// GroupEvents is a slice of GroupEvent.
type GroupEvents struct {
	GroupEvents []GroupEvent
}

// GroupEventAttendee is a struct that holds group event attendee data.
type GroupEventAttendee struct {
	EventId         string    `json:"event_id"`
	AttendeeId      string    `json:"attendee_id"`
	AttendingStatus int       `json:"attending_status"`
	CreationDate    time.Time `json:"creation_date"`
}

// GroupEventAttendees is a slice of GroupEventAttendee.
type GroupEventAttendees struct {
	GroupEventAttendees []GroupEventAttendee
}
