package helpers

import "github.com/gofrs/uuid"

/*
CreateUUID creates a new Universally Unique Identifier (UUID).

The function takes no parameters and returns a randomly generated NewV4 UUID
with a length of 36 cahracters and must contain four hyphens.

returns:
  - string: a randomly generated UUID as a string
  - error: an error if one occurs when generating the UUID.

Example:
  - The function will be used when creating a user id.
*/
func CreateUUID() (string, error) {
	newUUID, err := uuid.NewV4()
	return newUUID.String(), err
}
