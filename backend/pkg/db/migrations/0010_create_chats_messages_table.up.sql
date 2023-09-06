CREATE TABLE Chats_Messages (
  message_id TEXT NOT NULL PRIMARY KEY,
  chat_id TEXT NOT NULL,
  sender_id TEXT NOT NULL,
  message TEXT NOT NULL,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(chat_id) REFERENCES Chats(chat_id),
  FOREIGN KEY(sender_id) REFERENCES Users(user_id)
);