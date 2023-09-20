package dbstatements

import (
	"database/sql"
	"fmt"
)

func initGroupDBStatements(db *sql.DB) error {
	var err error

	/*INSERT*/

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
		request_pending,
		permission_level
	) VALUES (
		?, ?, ?, ?
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
		event_start_time
	) VALUES (
		?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups events statement: %w", err)
	}

	InsertGroupsEventsAttendees, err = db.Prepare(`
	INSERT INTO Groups_Events_Attendees (
		event_id,
		attendee_id,
		attending_status
	) VALUES (
		?, ?, ?
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert groups events attendees statement: %w", err)
	}

	/*SELECT*/

	SelectAllGroupsStmt, err = db.Prepare(`
	SELECT * FROM Groups
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all groups statement: %w", err)
	}

	SelectGroupByTitleStmt, err = db.Prepare(`
	SELECT * FROM Groups
	WHERE title = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select group statement: %w", err)
	}

	SelectGroupByIDStmt, err = db.Prepare(`
	SELECT * FROM Groups
	WHERE group_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select group statement: %w", err)
	}

	SelectAllGroupMembersStmt, err = db.Prepare(`
	SELECT * FROM Groups_Members
	WHERE group_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all group members statement: %w", err)
	}

	SelectUserGroupsStmt, err = db.Prepare(`
	SELECT G.*
	FROM Groups G
	INNER JOIN Groups_Members GM ON G.group_id = GM.group_id
	WHERE GM.member_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all group members statement: %w", err)
	}

	SelectAllGroupEventsStmt, err = db.Prepare(`
	SELECT * FROM Groups_Events
	WHERE group_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all group events statement: %w", err)
	}

	SelectAllEventAttendeesStmt, err = db.Prepare(`
	SELECT * FROM Groups_Events_Attendees
	WHERE event_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare select all group events attendees statement: %w", err)
	}

	/*UPDATE*/

	UpdateGroupMemberStatusStmt, err = db.Prepare(`
	UPDATE Groups_Members
	SET request_pending = 0, permission_level = ?
	WHERE member_id = ? AND group_id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare update group member status statement: %w", err)
	}

	UpdateAttendeeStatus, err = db.Prepare(`
	UPDATE Groups_Events_Attendees
	SET attending_status = ?
	WHERE attendee_id = ? AND event_id = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare update group event attendee status statement: %w", err)
	}

	return nil
}
