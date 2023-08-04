package dbstatements

import (
	"database/sql"
	"fmt"
)

const ()

// Prepared insert statements
var (
	InsertUserStmt                  *sql.Stmt
	InsertSessionsStmt              *sql.Stmt
	InsertPostStmt                  *sql.Stmt
	InsertPostsSelectedFollowerStmt *sql.Stmt
	InsertCommentsStmt              *sql.Stmt
	InsertReactionsStmt             *sql.Stmt
	InsertChatsStmt                 *sql.Stmt
	InsertChatsMessagesStmt         *sql.Stmt
	InsertFollowersStmt             *sql.Stmt
	InsertGroupsStmt                *sql.Stmt
	InsertGroupsMembersStmt         *sql.Stmt
	InsertGroupsEventsStmt          *sql.Stmt
	InsertGroupsEventsAttendees     *sql.Stmt
	InsertNotificationsStmt         *sql.Stmt
)

func InitDBStatements(db *sql.DB) error {
	var err error
	InsertUserStmt, err = db.Prepare(`
	INSERT INTO Users (
		user_id,
		is_logged_in,
		email,
		hashed_password,
		first_name,
		last_name, 
		date_of_birth,
		avatar_path,
		display_name,
		about_me
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert users statement: %s", err)
	}

	InsertSessionsStmt, err = db.Prepare(`
	INSERT INTO Sessions (
		session_id,
		user_id
	) VALUES (
		?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert sessions statement: %s", err)
	}

	InsertPostStmt, err = db.Prepare(`
	INSERT INTO Posts (
		post_id,
		group_id,
		creator_id,
		title,
		image_path,
		content,
		num_of_comments,
		privacy_level,
		likes,
		dislikes
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert posts statement: %s", err)
	}

	InsertPostsSelectedFollowerStmt, err = db.Prepare(`
	INSERT INTO Posts_Selected_Followers  (
		post_id,
		allowed_follower_id
	) VALUES (
		?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert post follower statement: %s", err)
	}

	InsertCommentsStmt, err = db.Prepare(`
	INSERT INTO Comments (
		comment_id,
		user_id,
		post_id,
		content,
		image_path,
		likes,
		dislikes
	) VALUES (
		?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert comments statement: %s", err)
	}

	InsertReactionsStmt, err = db.Prepare(`
	INSERT INTO Reactions (
		user_id,
		post_id,
		comment_id,
		reaction
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert reactions statement: %s", err)
	}

	InsertChatsStmt, err = db.Prepare(`
	INSERT INTO Chats (
		chat_id,
		sender_id,
		receiver_id
	) VALUES (
		?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert chats statement: %s", err)
	}

	InsertChatsMessagesStmt, err = db.Prepare(`
	INSERT INTO Chats_Messages (
		message_id,
		chat_id,
		sender_id,
		message
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert chats messages statement: %s", err)
	}

	InsertFollowersStmt, err = db.Prepare(`
	INSERT INTO Followers (
		follower_id,
		followee_id,
		following_status,
		request_pending
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert followers statement: %s", err)
	}

	InsertGroupsStmt, err = db.Prepare(`
	INSERT INTO Groups (
		group_id,
		title,
		description,
		creator_id
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups statement: %s", err)
	}

	InsertGroupsMembersStmt, err = db.Prepare(`
	INSERT INTO Groups_Members (
		group_id,
		member_id,
		request_pending
	) VALUES (
		?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups members statement: %s", err)
	}

	InsertGroupsEventsStmt, err = db.Prepare(`
	INSERT INTO Groups_Events (
		event_id,
		group_id,
		creator_id,
		title,
		description,
		event_start_time,
		total_going,
		total_not_going
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups events statement: %s", err)
	}

	InsertGroupsEventsAttendees, err = db.Prepare(`
	INSERT INTO Groups_Events_Attendees (
		event_id,
		attendee_id,
		attending_status,
		event_status
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups events attendees statement: %s", err)
	}

	InsertNotificationsStmt, err = db.Prepare(`
	INSERT INTO Notifications (
		notification_id,
		sender_id,
		receiver_id,
		group_id,
		post_id,
		event_id,
		comment_id,
		chat_id,
		reaction_type,
		read_status
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert notifications statement: %s", err)
	}

	return nil
}

func CloseDBStatements() {
	InsertUserStmt.Close()
	InsertSessionsStmt.Close()
	InsertPostStmt.Close()
	InsertPostsSelectedFollowerStmt.Close()
	InsertCommentsStmt.Close()
	InsertReactionsStmt.Close()
	InsertChatsStmt.Close()
	InsertChatsMessagesStmt.Close()
	InsertFollowersStmt.Close()
	InsertGroupsStmt.Close()
	InsertGroupsMembersStmt.Close()
	InsertGroupsEventsStmt.Close()
	InsertGroupsEventsAttendees.Close()
	InsertNotificationsStmt.Close()
}
