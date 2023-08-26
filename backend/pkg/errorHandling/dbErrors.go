package errorhandling

import "errors"

var (
	ErrNoRowsAffected    = errors.New("no rows affected")
	ErrMissingProfilePic = errors.New("missing or no profile pic")
)
