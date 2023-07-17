package helpers_test

import (
	db "socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/helpermodels"
	"testing"
)

func TestIsValidPath(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"../helpers", true},
		{"./helpers01", false},
		{"../db/migrations", true},
		{"../models/migrations", false},
	}
	for _, c := range cases {
		got, err := helpers.IsValidPath(c.in)
		if err != nil {
			t.Log(err)
		}
		if got != c.want {
			t.Errorf("IsValidPath(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCheckValidePath(t *testing.T) {
	tempPath := t.TempDir()
	// log.Println(tempPath)

	testCases := []struct {
		caseName         string
		fileName         string
		directory        string
		extension        string
		isAlphaNumeric   bool
		isValidDirectory bool
		isNotDuplicate   bool
		expectError      bool
	}{
		{
			caseName:         "All is working",
			fileName:         "mydatabase",
			directory:        tempPath,
			isAlphaNumeric:   true,
			isValidDirectory: true,
			isNotDuplicate:   true,
			expectError:      false,
		},
		{
			caseName:         "Isn't alphanumeric",
			fileName:         "#fail1",
			directory:        tempPath,
			isAlphaNumeric:   false,
			isValidDirectory: false,
			isNotDuplicate:   true,
			expectError:      true,
		},
		{
			caseName:         "Isn't a valid directory",
			fileName:         "fail2",
			directory:        "/inexistant/Path",
			isAlphaNumeric:   true,
			isValidDirectory: false,
			isNotDuplicate:   true,
			expectError:      true,
		},
		{
			caseName:         "Is a duplicate",
			fileName:         "fail3",
			directory:        tempPath,
			isAlphaNumeric:   true,
			isValidDirectory: true,
			isNotDuplicate:   false,
			expectError:      true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			// Perform the test
			dbFilePath := &helpermodels.FilePathComponents{
				FileName:  tc.fileName,
				Directory: tc.directory,
				Extension: ".db",
			}

			_, err := helpers.CheckValidPath(dbFilePath)

			if !tc.isNotDuplicate {
				db.CreateDatabase(dbFilePath)
				_, err = helpers.CheckValidPath(dbFilePath)
			}

			if tc.expectError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

		})
	}
}
