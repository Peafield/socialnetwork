package dbutils_test

import (
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/models/helpermodels"
	"testing"
)

func TestOpenDatabase(t *testing.T) {
	testCases := []struct {
		Name          string
		FilePath      *helpermodels.FilePathComponents
		ExpectedError bool
	}{
		{
			Name: "Valid path",
			FilePath: &helpermodels.FilePathComponents{
				Directory: "../",
				FileName:  "testDB",
				Extension: ".db",
			},
			ExpectedError: false,
		},
		{
			Name: "Invalid path",
			FilePath: &helpermodels.FilePathComponents{
				Directory: "/invalid/path/",
				FileName:  "testDB",
				Extension: ".db",
			},
			ExpectedError: true,
		},
		{
			Name: "Invalid filename",
			FilePath: &helpermodels.FilePathComponents{
				Directory: "../",
				FileName:  "mockTest",
				Extension: ".db",
			},
			ExpectedError: true,
		},
		{
			Name: "Invalid extension",
			FilePath: &helpermodels.FilePathComponents{
				Directory: "../",
				FileName:  "testDB",
				Extension: ".txt",
			},
			ExpectedError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := dbutils.OpenDatabase(tc.FilePath)
			defer dbutils.CloseDatabase()
			if tc.ExpectedError {
				if err == nil {
					t.Errorf("expected and error but did not get one")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but go one: %s", err)
				}
			}

			err = dbutils.DB.Ping()
			if tc.ExpectedError {
				if err == nil {
					t.Errorf("expected and error but did not get one")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but go one: %s", err)
				}
			}
		})
	}
}
