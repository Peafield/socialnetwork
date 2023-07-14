CREATE TABLE Chats (
  chat_id TEXT NOTNULL PRIMARY KEY,
  sender_id TEXT NOTNULL,
  receiver_id TEXT NOTNULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(sender_id) REFERENCES Users(user_id),
  FOREIGN KEY(receiver_id) REFERENCES Users(user_id)
);
