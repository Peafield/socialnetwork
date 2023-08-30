CREATE TABLE Users (
  user_id TEXT UNIQUE NOT NULL PRIMARY KEY,
  is_logged_in INTEGER NOT NULL DEFAULT 0,
  is_private INTEGER NOT NULL DEFAULT 0,
  email TEXT UNIQUE NOT NULL,
  display_name TEXT UNIQUE,
  hashed_password TEXT NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  avatar_path TEXT NOT NULL,
  about_me TEXT,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);