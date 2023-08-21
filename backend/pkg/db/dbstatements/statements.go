package dbstatements

import (
	"database/sql"
	"fmt"
)

// Prepared insert statements
var (
	/*Insert Statements*/
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

	/*Update Statments*/
	UpdatePostNumOfComments   *sql.Stmt
	UpdateAllUsersToSignedOut *sql.Stmt
	UpdateUserLoggedIn        *sql.Stmt
	UpdateUserLoggedOut       *sql.Stmt

	/*Select Statements*/
	SelectUserByID          string
	SelectUserByDisplayName string
	SelectUserViewablePosts string
	SpecificUserPosts       string
)

func InitDBStatements(db *sql.DB) error {
	var err error

	err = initUserDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare user statements: %w", err)
	}

	err = initPostDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare post statements: %w", err)
	}

	err = initCommentDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare comment statements: %w", err)
	}

	InsertSessionsStmt, err = db.Prepare(`
	INSERT INTO Sessions (
		session_id,
		user_id
	) VALUES (
		?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert sessions statement: %w", err)
	}

	InsertPostsSelectedFollowerStmt, err = db.Prepare(`
	INSERT INTO Posts_Selected_Followers  (
		post_id,
		allowed_follower_id
	) VALUES (
		?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert post follower statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert reactions statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert chats statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert chats messages statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert followers statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert groups statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert groups members statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert groups events statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert groups events attendees statement: %w", err)
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
		return fmt.Errorf("failed to prepare insert notifications statement: %w", err)
	}

	return nil
}

func CloseDBStatements() {
	/*Insert Statment Closure*/
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

	/*Update Statement Closure*/
	UpdatePostNumOfComments.Close()
	UpdateAllUsersToSignedOut.Close()
	UpdateUserLoggedIn.Close()
	UpdateUserLoggedOut.Close()
}
