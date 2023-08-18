CREATE TABLE Groups_Members (
  group_id TEXT NOT NULL,
  member_id TEXT NOT NULL,
  request_pending INT,
  permission_level tinyint DEFAULT 0 CHECK (permission_level <= 2), 
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(member_id) REFERENCES Users(user_id)
);