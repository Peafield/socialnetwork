package helpers

import (
	"fmt"
	"socialnetwork/pkg/models/helpermodels"
)

func FileCreator(fullPath string, fileCreator helpermodels.FileCreator) error {
	file, err := fileCreator.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file path: %s", err)
	}
	defer file.Close()
	return nil
}
