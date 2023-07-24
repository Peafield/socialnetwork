package helpers

import "time"

// NormalizeTime converts a time.Time value to UTC and then back to the system's local time zone.
func NormalizeTime(t time.Time) time.Time {
	utc := t.UTC()
	return utc.Local()
}
