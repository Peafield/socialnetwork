CREATE TABLE Groups_Events (
-- PRIMARY KEY = NOT NULL + UNIQUE
  event_id TEXT PRIMARY KEY,
  group_id TEXT NOT NULL,
  creator_id TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  event_start_time DATE NOT NULL,
  total_going INTEGER DEFAULT 0,
  total_not_going INTEGER DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);