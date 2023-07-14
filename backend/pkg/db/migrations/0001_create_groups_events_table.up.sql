CREATE TABLE Groups_Events (
  event_id TEXT UNIQUE NOTNULL PRIMARY KEY,
  group_id TEXT NOTNULL,
  creator_id TEXT NOTNULL,
  title TEXT NOTNULL,
  description TEXT NOTNULL,
  event_start_time TIMESTAMP,
  total_going INTEGER DEFAULT 0,
  total_not_going INTEGER DEFAULT 0,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);