CREATE TABLE Reactions (
    user_id TEXT NOT NULL,
    post_id TEXT,
    comment_id TEXT,
    reaction TEXT,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES Users(user_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id),
    FOREIGN KEY(comment_id) REFERENCES Comments(comment_id)
);