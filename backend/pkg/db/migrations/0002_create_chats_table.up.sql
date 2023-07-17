CREATE TABLE Chats (
  chat_id TEXT NOT NULL PRIMARY KEY,
  sender_id TEXT NOT NULL,
  receiver_id TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(sender_id) REFERENCES Users(user_id),
  FOREIGN KEY(receiver_id) REFERENCES Users(user_id)
);
