package helpers

import (
	"os"
)

/*
IsValidPath validates a given path.

It uses os.Stat method to check whether the directory exists,
if not it will return false and prompt an error.

Parameters:
  - path: a path to a directory to be tested as a string.

Returns:
  - bool: true or false depending on the match.
  - error: if an error occurs during the matching.

Example:
  - Testing the directory file path for file storage (e.g databases, images, etc)
*/
func IsValidPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}
