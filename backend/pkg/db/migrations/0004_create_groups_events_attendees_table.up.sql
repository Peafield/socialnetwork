CREATE TABLE Groups_Events_Attendees (
  event_id TEXT NOT NULL PRIMARY KEY,
  attendee_id TEXT NOT NULL,
  attending_status TEXT NOT NULL,
  event_status INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(event_id) REFERENCES Groups_Events(event_id),
  FOREIGN KEY(attendee_id) REFERENCES Users(user_id)
);