package errorhandling

import "errors"

var (
	ErrNoRowsAffected = errors.New("no rows affected")
)
