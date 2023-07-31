CREATE TABLE Followers (
  follower_id TEXT NOT NULL,
  followee_id TEXT NOT NULL,
  following_status INTEGER NOT NULL DEFAULT 0,
  request_pending INTEGER NOT NULL DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(follower_id) REFERENCES Users(user_id),
  FOREIGN KEY(followee_id) REFERENCES Users(user_id)
);