CREATE TABLE Posts (
  post_id TEXT NOTNULL PRIMARY KEY,
  group_id TEXT,
  creator_id TEXT NOTNULL,
  title TEXT NOTNULL,
  image_path TEXT,
  content TEXT NOTNULL,
  privacy_level INTEGER NOTNULL,
  allowed_followers TEXT,
  likes INTEGER DEFAULT 0,
  dislikes INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(creator_id) REFERENCES Users(user_id),
  FOREIGN KEY(group_id) REFERENCES Groups(group_id)
);