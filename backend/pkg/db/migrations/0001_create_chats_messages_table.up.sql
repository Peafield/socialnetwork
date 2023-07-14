CREATE TABLE Chats_Messages (
  message_id TEXT NOTNULL PRIMARY KEY,
  chat_id TEXT NOTNULL,
  sender_id TEXT NOTNULL,
  message TEXT NOTNULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(chat_id) REFERENCES Chats(chat_id),
  FOREIGN KEY(sender_id) REFERENCES Users(user_id)
);