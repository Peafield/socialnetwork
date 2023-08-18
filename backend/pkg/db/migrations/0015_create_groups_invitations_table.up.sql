CREATE TABLE Groups_Invitations (
  group_id TEXT NOT NULL, 
  user_id TEXT NOT NULL, 
  is_invited INT, 
  FOREIGN KEY(group_id) REFERENCES Groups(group_id), 
  FOREIGN KEY(user_id) REFERENCES Users(user_id)
);