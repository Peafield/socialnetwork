package helpers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"socialnetwork/pkg/models/helpermodels"
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
	if _, err := os.Stat(path); err == nil {
		return true, err
	} else if errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	return true, nil
}

/*
CheckValidPath validates the format and the existence of a given path.
The function returns true and nil if the following conditions are satisfied:

  - Must be alphanumeric
  - Must provide a valid directory
  - Must be unique within the specified directory

Otherwise, the function returns false along its corresponding error.

Parameters:
  - Filepath (FilePathManager): An interface which provides information related to the ressource (e.g: file name, directory and/or file extension) .

Returns:
  - bool
  - error
*/
func CheckValidPath(filePath helpermodels.FilePathManager) (bool, error) {
	fileName := filePath.GetFileName()
	directory := filePath.GetDirectory()
	extension := filePath.GetFileExtension()

	isFileNameValid, err := IsAlphaNumeric(fileName)
	if !isFileNameValid {
		return false, fmt.Errorf("file name contains non alpha-numeric characters. Err: %s", err)
	}

	isValidDirPath, err := IsValidPath(directory)

	if !isValidDirPath {
		return false, fmt.Errorf("directory is not valid. Err: %s", err)
	}

	fullPath := path.Join(directory, fileName+extension)
	log.Println(fullPath)
	fullPathExists, err := IsValidPath(fullPath)

	if fullPathExists && err == nil {
		return false, fmt.Errorf("file path already exists: %s", err)
	}
	return true, nil
}
