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
	InsertGroupChatStmt             *sql.Stmt

	/*Update Statments*/
	UpdatePostImagePathStmt           *sql.Stmt
	UpdatePostContentStmt             *sql.Stmt
	UpdatePostPrivacyLevelStmt        *sql.Stmt
	UpdatePostIncreaseNumOfComments   *sql.Stmt
	UpdatePostDecreaseNumOfComments   *sql.Stmt
	UpdateAllUsersToSignedOut         *sql.Stmt
	UpdateUserLoggedIn                *sql.Stmt
	UpdateUserLoggedOut               *sql.Stmt
	UpdateUserAccountStmt             *sql.Stmt
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
	UpdateAllUserNotifications        *sql.Stmt
	UpdateGroupMemberStatusStmt       *sql.Stmt
	UpdateAttendeeStatus              *sql.Stmt

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
	SelectFolloweesOfUserStmt                    *sql.Stmt
	SelectFollowersOfUserStmt                    *sql.Stmt
	SelectChatMessagesByChatIdStmt               *sql.Stmt
	SelectChatBySenderAndRecieverIdStmt          *sql.Stmt
	SelectAllChatsByUserIdStmt                   *sql.Stmt
	SelectAllUserNotifications                   *sql.Stmt
	SelectAllGroupsStmt                          *sql.Stmt
	SelectGroupByTitleStmt                       *sql.Stmt
	SelectGroupByIDStmt                          *sql.Stmt
	SelectAllGroupMembersStmt                    *sql.Stmt
	SelectUserGroupsStmt                         *sql.Stmt
	SelectGroupPostsStmt                         *sql.Stmt
	SelectAllGroupEventsStmt                     *sql.Stmt
	SelectAllEventAttendeesStmt                  *sql.Stmt
	SelectGroupChatStmt                          *sql.Stmt

	/*Delete Statements*/
	DeleteUserAccountStmt           *sql.Stmt
	DeletePostStmt                  *sql.Stmt
	DeleteUserComment               *sql.Stmt
	DeletePostReaction              *sql.Stmt
	DeleteCommentReaction           *sql.Stmt
	DeleteFollowerStmt              *sql.Stmt
	DeletePostsSelectedFollowerStmt *sql.Stmt
	DeleteNotificationStmt          *sql.Stmt
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

	err = initChatDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare chat statements: %w", err)
	}

	err = initNotificationStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare notification statements: %w", err)
	}

	err = initGroupDBStatements(db)
	if err != nil {
		return fmt.Errorf("failed to prepare group statements: %w", err)
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
	InsertGroupChatStmt.Close()

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
	SelectFolloweesOfUserStmt.Close()
	SelectFollowersOfUserStmt.Close()
	SelectChatMessagesByChatIdStmt.Close()
	SelectChatBySenderAndRecieverIdStmt.Close()
	SelectAllChatsByUserIdStmt.Close()
	SelectAllUserNotifications.Close()
	SelectAllGroupsStmt.Close()
	SelectGroupByTitleStmt.Close()
	SelectGroupByIDStmt.Close()
	SelectAllGroupMembersStmt.Close()
	SelectUserGroupsStmt.Close()
	SelectGroupPostsStmt.Close()
	SelectAllGroupEventsStmt.Close()
	SelectAllEventAttendeesStmt.Close()
	SelectGroupChatStmt.Close()

	/*Update Statement Closure*/
	UpdatePostImagePathStmt.Close()
	UpdatePostContentStmt.Close()
	UpdatePostPrivacyLevelStmt.Close()
	UpdatePostIncreaseNumOfComments.Close()
	UpdatePostDecreaseNumOfComments.Close()
	UpdateAllUsersToSignedOut.Close()
	UpdateUserAccountStmt.Close()
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
	UpdateAllUserNotifications.Close()
	UpdateGroupMemberStatusStmt.Close()
	UpdateAttendeeStatus.Close()

	/*Delete Statement Closure*/
	DeleteUserAccountStmt.Close()
	DeletePostStmt.Close()
	DeleteUserComment.Close()
	DeletePostReaction.Close()
	DeleteCommentReaction.Close()
	DeleteFollowerStmt.Close()
	DeletePostsSelectedFollowerStmt.Close()
	DeleteNotificationStmt.Close()
}
