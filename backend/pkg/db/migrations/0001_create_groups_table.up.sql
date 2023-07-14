CREATE TABLE Groups (
  group_id TEXT UNIQUE NOTNULL PRIMARY KEY,
  title TEXT NOTNULL,
  description TEXT,
  creator_id TEXT NOTNULL,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);