package helpers

import "regexp"

/*
IsAlphaNumeric takes a string as input and checks with all the characters
in the string are alphanumeric.

It uses a regex pattern that will only except alphanumeric values and matches these
against the string. It will return a bool or true there are only alphanumeric characters
and false if not. It will also return an error if there is one.

Parameters:
  - input: a string to be tested.

Returns:
  - bool: true or false depending on the match.
  - error: if an error occurs during the matching.

Example:
  - Used when checking database names or checking usernames.
*/
func IsAlphaNumeric(input string) (bool, error) {
	pattern := `^[a-zA-Z0-9]*$`
	match, err := regexp.MatchString(pattern, input)
	return match, err
}
