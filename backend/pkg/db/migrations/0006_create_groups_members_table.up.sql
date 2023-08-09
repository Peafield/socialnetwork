CREATE TABLE Groups_Members (
  group_id TEXT NOT NULL,
  member_id TEXT NOT NULL,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(member_id) REFERENCES Users(user_id)
);