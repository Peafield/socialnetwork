CREATE TABLE Followers (
  followee_id TEXT NOT NULL PRIMARY KEY,
  follower_id TEXT NOT NULL,
  following_status INTEGER NOT NULL DEFAULT 0,
  request_pending INTEGER NOT NULL DEFAULT 0,
  followed_at_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(followee_id) REFERENCES Users(user_id),
  FOREIGN KEY(follower_id) REFERENCES Users(user_id)
);