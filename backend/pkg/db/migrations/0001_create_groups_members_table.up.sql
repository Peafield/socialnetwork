CREATE TABLE Groups_Members (
  group_id TEXT NOTNULL,
  member_id TEXT NOTNULL,
  request_pending INTEGER NOTNULL DEFAULT 0,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(member_id) REFERENCES Users(user_id)
);