CREATE TABLE Groups (
  group_id TEXT UNIQUE NOT NULL PRIMARY KEY,
  title TEXT NOT NULL UNIQUE,
  description TEXT,
  creator_id TEXT NOT NULL,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id)
);