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
	UpdatePostImagePathStmt           *sql.Stmt
	UpdatePostContentStmt             *sql.Stmt
	UpdatePostPrivacyLevelStmt        *sql.Stmt
	UpdatePostIncreaseNumOfComments   *sql.Stmt
	UpdatePostDecreaseNumOfComments   *sql.Stmt
	UpdateAllUsersToSignedOut         *sql.Stmt
	UpdateUserLoggedIn                *sql.Stmt
	UpdateUserLoggedOut               *sql.Stmt
	UpdatePostReaction                *sql.Stmt
	UpdateCommentReaction             *sql.Stmt
	UpdateCommentContent              *sql.Stmt
	UpdateCommentImagePath            *sql.Stmt
	UpdatePostIncreaseLikeStmt        *sql.Stmt
	UpdatePostDecreaseLikesStmt       *sql.Stmt
	UpdatePostIncreaseDislikeStmt     *sql.Stmt
	UpdatePostDecreaseDislikesStmt    *sql.Stmt
	UpdateCommentIncreaseLikeStmt     *sql.Stmt
	UpdateCommentDecreaseLikesStmt    *sql.Stmt
	UpdateCommentIncreaseDislikeStmt  *sql.Stmt
	UpdateCommentDecreaseDislikesStmt *sql.Stmt
	UpdateFollowingStatusStmt         *sql.Stmt
	UpdateRequestPendingStmt          *sql.Stmt

	/*Select Statements*/
	SelectAllUsersStmt                           *sql.Stmt
	SelectUserByIDStmt                           *sql.Stmt
	SelectUserByDisplayNameStmt                  *sql.Stmt
	SelectUserByEmailAndDisplayNameAndUserIdStmt *sql.Stmt
	SelectUserByEmailOrDisplayNameStmt           *sql.Stmt
	SelectUserByIDOrDisplayNameStmt              *sql.Stmt
	SelectAllPostsStmt                           *sql.Stmt
	SelectUserViewablePostsStmt                  *sql.Stmt
	SelectSpecificUserPostsStmt                  *sql.Stmt
	SelectPostCommentsStmt                       *sql.Stmt
	SelectPostReactionStmt                       *sql.Stmt
	SelectCommentReactionStmt                    *sql.Stmt
	SelectFollowerInfoStmt                       *sql.Stmt

	/*Delete Statements*/
	DeleteUserAccountStmt           *sql.Stmt
	DeletePostStmt                  *sql.Stmt
	DeleteUserComment               *sql.Stmt
	DeletePostReaction              *sql.Stmt
	DeleteCommentReaction           *sql.Stmt
	DeleteFollowerStmt              *sql.Stmt
	DeletePostsSelectedFollowerStmt *sql.Stmt
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

	err = initReactionDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare reaction statements: %w", err)
	}

	err = initFollowerDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare follower statements: %w", err)
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
		reaction_type
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?
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

	/*Select Statment Closure*/
	SelectAllUsersStmt.Close()
	SelectUserByIDStmt.Close()
	SelectUserByDisplayNameStmt.Close()
	SelectUserByEmailAndDisplayNameAndUserIdStmt.Close()
	SelectUserByIDOrDisplayNameStmt.Close()
	SelectUserByEmailOrDisplayNameStmt.Close()
	SelectAllPostsStmt.Close()
	SelectUserViewablePostsStmt.Close()
	SelectSpecificUserPostsStmt.Close()
	SelectPostReactionStmt.Close()
	SelectCommentReactionStmt.Close()
	SelectPostCommentsStmt.Close()
	SelectFollowerInfoStmt.Close()

	/*Update Statement Closure*/
	UpdatePostImagePathStmt.Close()
	UpdatePostContentStmt.Close()
	UpdatePostPrivacyLevelStmt.Close()
	UpdatePostIncreaseNumOfComments.Close()
	UpdatePostDecreaseNumOfComments.Close()
	UpdateAllUsersToSignedOut.Close()
	UpdateUserLoggedIn.Close()
	UpdateUserLoggedOut.Close()
	UpdatePostReaction.Close()
	UpdateCommentReaction.Close()
	UpdateCommentContent.Close()
	UpdateCommentImagePath.Close()
	UpdatePostIncreaseLikeStmt.Close()
	UpdatePostDecreaseLikesStmt.Close()
	UpdatePostIncreaseDislikeStmt.Close()
	UpdatePostDecreaseDislikesStmt.Close()
	UpdateCommentIncreaseLikeStmt.Close()
	UpdateCommentDecreaseLikesStmt.Close()
	UpdateCommentIncreaseDislikeStmt.Close()
	UpdateCommentDecreaseDislikesStmt.Close()
	UpdateFollowingStatusStmt.Close()
	UpdateRequestPendingStmt.Close()

	/*Delete Statement Closure*/
	DeleteUserAccountStmt.Close()
	DeletePostStmt.Close()
	DeleteUserComment.Close()
	DeletePostReaction.Close()
	DeleteCommentReaction.Close()
	DeleteFollowerStmt.Close()
	DeletePostsSelectedFollowerStmt.Close()
}
