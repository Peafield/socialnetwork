package errorhandling

import "errors"

var (
	ErrNoRowsAffected    = errors.New("no rows affected")
	ErrNoResultsFound    = errors.New("no results found")
	ErrMissingProfilePic = errors.New("missing or no profile pic")
	ErrUserExists        = errors.New("user display name or email already in use")
)
