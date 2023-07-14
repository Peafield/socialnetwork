CREATE TABLE Followers (
  followee_id TEXT NOTNULL PRIMARY KEY,
  follower_id TEXT NOTNULL,
  following_status INTEGER NOTNULL DEFAULT 0,
  request_pending INTEGER NOTNULL DEFAULT 0,
  followed_at_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(followee) REFERENCES Users(user_id),
  FOREIGN KEY(follower) REFERENCES Users(user_id)
);