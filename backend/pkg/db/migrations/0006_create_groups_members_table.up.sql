CREATE TABLE Groups_Members (
  group_id TEXT NOT NULL,
  member_id TEXT NOT NULL,
  request_pending INTEGER NOT NULL DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(member_id) REFERENCES Users(user_id)
);