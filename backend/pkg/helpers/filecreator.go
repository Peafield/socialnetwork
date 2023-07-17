package helpers

import (
	"fmt"
	"socialnetwork/pkg/models/helpermodels"
)

/*
File Creator creates a file given a full file path (with extension).

Parameters:
  - fullPath: full path of a file including the extension.

Returns:
  - error: if an error occurs during the creating process.

Example:
  - Create files locally, including database creation, images, etc.
*/
func FileCreator(fullPath string, fileCreator helpermodels.FileCreator) error {
	file, err := fileCreator.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file path: %s", err)
	}
	defer file.Close()
	return nil
}
