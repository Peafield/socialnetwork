package helpers_test

import (
	"socialnetwork/pkg/helpers"
	"testing"
)

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		Name           string
		Password       string
		HashedPassword string
		ExpectedError  bool
	}{
		{
			Name:           "Valid Hashing",
			Password:       "password",
			HashedPassword: "$2a$10$CzqaN60QCVufP1lRnhWDBuRKR.Z4sS3q3aBAhu44mNInbkIUKrk8S",
			ExpectedError:  false,
		},
		{
			Name:           "Password as an empty string",
			Password:       "",
			HashedPassword: "$2a$10$CzqaN60QCVufP1lRnhWDBuRKR.Z4sS3q3aBAhu44mNInbkIUKrk8S",
			ExpectedError:  true,
		},

		{
			Name:           "Invalid Hashing",
			Password:       "password",
			HashedPassword: "$2a$10$iRS.fdyfl/lXAqzx7Q3zzO9zGFc20BiGXFQl149P33hDBWAp1yXvC",
			ExpectedError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			hashedPassword, err := helpers.HashPassword(tc.Password)
			if hashedPassword == tc.HashedPassword && tc.ExpectedError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if hashedPassword != tc.HashedPassword && !tc.ExpectedError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}

func TestCompareHashedPassword(t *testing.T) {
	testCases := []struct {
		Name           string
		Password       string
		HashedPassword string
		ExpectedError  bool
	}{
		{
			Name:           "Valid Password",
			Password:       "password",
			HashedPassword: "$2a$10$CzqaN60QCVufP1lRnhWDBuRKR.Z4sS3q3aBAhu44mNInbkIUKrk8S",
			ExpectedError:  false,
		},
		{
			Name:           "Invalid Password",
			Password:       "password",
			HashedPassword: "$2a$10$iRS.fdyfl/lXAqzx7Q3zzO9zGFc20BiGXFQl149P33hDBWAp1yXvC",
			ExpectedError:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := helpers.CompareHashedPassword(tc.HashedPassword, tc.Password)
			if tc.ExpectedError && err == nil {
				t.Error("Expected an error, but got nil")
			} else if !tc.ExpectedError && err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
		})
	}
}
