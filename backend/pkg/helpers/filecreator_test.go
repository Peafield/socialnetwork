package helpers_test

import (
	"errors"
	"os"
	"path"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/helpermodels"
	"testing"
)

type MockFileCreator struct{}

func (m *MockFileCreator) Create(name string) (*os.File, error) {
	return nil, errors.New("Failed to create file in given directory")
}
func TestFileCreator(t *testing.T) {
	tempPath := t.TempDir()

	testCases := []struct {
		caseName    string
		fullPath    string
		fileCreator helpermodels.FileCreator
		isCreated   bool
		expectError bool
	}{
		{
			caseName:    "File created, .db extension",
			fullPath:    path.Join(tempPath, "myfile.db"),
			fileCreator: &helpermodels.OSFileCreator{},
			isCreated:   true,
			expectError: false,
		},
		{
			caseName:    "File created, .txt extension",
			fullPath:    path.Join(tempPath, "myfile.txt"),
			fileCreator: &helpermodels.OSFileCreator{},
			isCreated:   true,
			expectError: false,
		},
		{
			caseName:    "File created, .jpeg extension",
			fullPath:    path.Join(tempPath, "myfile.txt"),
			fileCreator: &helpermodels.OSFileCreator{},
			isCreated:   true,
			expectError: false,
		},
		{
			caseName:    "Failed Attempt to creating a file",
			fullPath:    path.Join(tempPath, "myfile.db"),
			fileCreator: &MockFileCreator{},
			isCreated:   false,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			filepath := tc.fullPath

			err := helpers.FileCreator(filepath, tc.fileCreator)
			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})

	}
}
