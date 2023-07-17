CREATE TABLE Groups_Events (
  event_id TEXT UNIQUE NOT NULL PRIMARY KEY,
  group_id TEXT NOT NULL,
  creator_id TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  event_start_time TIMESTAMP,
  total_going INTEGER DEFAULT 0,
  total_not_going INTEGER DEFAULT 0,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);