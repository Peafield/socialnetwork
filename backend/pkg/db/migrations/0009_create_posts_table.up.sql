CREATE TABLE Posts (
  post_id TEXT NOT NULL PRIMARY KEY,
  group_id TEXT,
  creator_id TEXT NOT NULL,
  title TEXT NOT NULL,
  image_path TEXT,
  content TEXT NOT NULL,
  privacy_level INTEGER NOT NULL,
  allowed_followers TEXT,
  likes INTEGER DEFAULT 0,
  dislikes INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id),
  FOREIGN KEY(group_id) REFERENCES Groups(group_id)
);