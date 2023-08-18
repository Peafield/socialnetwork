CREATE TABLE Posts (
  post_id TEXT NOT NULL PRIMARY KEY,
  group_id TEXT DEFAULT '',
  creator_id TEXT NOT NULL,
  image_path TEXT,
  content TEXT NOT NULL,
  num_of_comments INTEGER DEFAULT 0,
  privacy_level INTEGER DEFAULT 0,
  likes INTEGER DEFAULT 0,
  dislikes INTEGER DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id),
  FOREIGN KEY(group_id) REFERENCES Groups(group_id)
);