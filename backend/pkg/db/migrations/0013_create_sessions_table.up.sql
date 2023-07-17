CREATE TABLE Sessions (
    session_id TEXT NOT NULL UNIQUE,
    user_id TEXT NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES Users(user_id)
);