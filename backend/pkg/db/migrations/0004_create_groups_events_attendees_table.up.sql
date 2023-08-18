CREATE TABLE Groups_Events_Attendees (
  event_id TEXT PRIMARY KEY,
  attendee_id TEXT NOT NULL,
  attending_status BOOL,
  event_status INTEGER DEFAULT 0,  --what does it mean?
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(event_id) REFERENCES Groups_Events(event_id),
  FOREIGN KEY(attendee_id) REFERENCES Users(user_id)
);