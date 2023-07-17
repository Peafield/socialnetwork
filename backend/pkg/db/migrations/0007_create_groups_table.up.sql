CREATE TABLE Groups (
  group_id TEXT UNIQUE NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  creator_id TEXT NOT NULL,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);