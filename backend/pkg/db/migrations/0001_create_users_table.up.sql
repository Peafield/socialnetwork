CREATE TABLE Users (
  user_id TEXT UNIQUE NOT NULL PRIMARY KEY,
  isLoggedIn INTEGER NOT NULL DEFAULT 0,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  avatar_path TEXT NOT NULL,
  nickname TEXT UNIQUE,
  about_me TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);