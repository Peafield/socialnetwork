package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword hashes a given password at the default cost and returns it.

It takes in a user inputted password and generates a hash of the password at the default
cost of 10.

Parameters:
  - password (string): an inputted password

Returns:
  - string: the hashed password
  - error: if one occurs during the hashing process

Examples:
  - The function will be used to initially hash the user's password upon registration.
*/
func HashPassword(password string) (string, error) {
	if len(password) < 1 {
		return "", fmt.Errorf("password is an empty string")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

/*
CompareHashedPassword compares a hashed password with an inputted password.

The function compares the hashed and inputted passwords and returns nil if successful or
an error if not.

Parameters:
  - hashedPassword (string): a hash of a password.
  - password (string): a plain text password.

Returns:
  - error: an error is returned if the hashed password and password are not the same, other it returns nil.

Example:
  - The function will be used when authenticating a user.
*/
func CompareHashedPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
