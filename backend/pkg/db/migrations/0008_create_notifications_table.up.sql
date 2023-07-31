
CREATE TABLE Notifications (
  notification_id TEXT NOT NULL PRIMARY KEY,
  sender_id TEXT,
  receiver_id TEXT,
  group_id TEXT,
  post_id TEXT,
  event_id TEXT,
  comment_id TEXT,
  chat_id TEXT,
  reaction_type TEXT,
  read_status INTEGER NOT NULL DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(sender_id) REFERENCES Users(user_id),
  FOREIGN KEY(receiver_id) REFERENCES Users(user_id),
  FOREIGN KEY(group_id) REFERENCES Groups(group_id),
  FOREIGN KEY(post_id) REFERENCES Posts(post_id),
  FOREIGN KEY(event_id) REFERENCES Groups_Events(event_id),
  FOREIGN KEY(comment_id) REFERENCES Comments(comment_id),
  FOREIGN KEY(chat_id) REFERENCES Chats(chat_id)
);