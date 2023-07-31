CREATE TABLE Posts_Selected_Followers (
  post_id TEXT NOT NULL,
  allowed_follower_id TEXT NOT NULL,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES Posts(post_id),
  FOREIGN KEY (allowed_follower_id) REFERENCES Users(user_id),
  UNIQUE(post_id, allowed_follower_id)
);