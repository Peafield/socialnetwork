package helpers_test

import (
	"reflect"
	"socialnetwork/pkg/helpers"
	"socialnetwork/pkg/models/dbmodels"
	"testing"
	"time"
)

var TEST_DOB = time.Date(1997, 2, 10, 7, 15, 30, 100, time.Local)
var TEST_CREATION_DATE = time.Date(2023, 3, 2, 15, 2, 30, 100, time.Local)
var TEST_CREATION_DATE_2 = time.Date(2023, 7, 23, 20, 15, 30, 100, time.Local)

type TestAddressStruct struct {
	Field1 string
	Field2 int
	Field3 bool
}

// makes sure the returned values are the same as the expected values, if the struct properties are not exported then the function panics are this is not viable using .Interface()
func TestStructFieldValues(t *testing.T) {
	testcases := []struct {
		name           string
		s              interface{}
		expectedValues []interface{}
		expectError    bool
	}{
		{
			name: "User struct example",
			s: &dbmodels.User{
				UserId:         "8519045",
				IsLoggedIn:     0,
				Email:          "example@gmail.com",
				HashedPassword: "webouf9259gAAAGhghgb",
				FirstName:      "Hazza",
				LastName:       "Gazza",
				DOB:            TEST_DOB,
				AvatarPath:     "/path/to/avatar.png",
				DisplayName:    "hazmnaz",
				AboutMe:        "My name is hazza, I am 26 and I work as a software engineer for Google making the latest and greatest AI super thing.",
				CreationDate:   TEST_CREATION_DATE,
			},
			expectedValues: []interface{}{
				"8519045",
				0,
				"example@gmail.com",
				"webouf9259gAAAGhghgb",
				"Hazza", "Gazza",
				TEST_DOB,
				"/path/to/avatar.png",
				"hazmnaz",
				"My name is hazza, I am 26 and I work as a software engineer for Google making the latest and greatest AI super thing.",
				TEST_CREATION_DATE,
			},
			expectError: false,
		},
		{
			name: "Not a pointer to a struct",
			s: dbmodels.User{
				UserId:         "8519045",
				IsLoggedIn:     0,
				Email:          "example@gmail.com",
				HashedPassword: "webouf9259gAAAGhghgb",
				FirstName:      "Hazza",
				LastName:       "Gazza",
				DOB:            time.Now().AddDate(-26, -5, -14),
				AvatarPath:     "/path/to/avatar.png",
				DisplayName:    "hazmnaz",
				AboutMe:        "My name is hazza, I am 26 and I work as a software engineer for Google making the latest and greatest AI super thing.",
				CreationDate:   time.Now().Add(-time.Hour * 5),
			},
			expectedValues: []interface{}{
				"8519045",
				0,
				"example@gmail.com",
				"webouf9259gAAAGhghgb",
				"Hazza", "Gazza",
				TEST_DOB,
				"/path/to/avatar.png",
				"hazmnaz",
				"My name is hazza, I am 26 and I work as a software engineer for Google making the latest and greatest AI super thing.",
				TEST_CREATION_DATE,
			},
			expectError: true,
		},
		{
			name: "Simple struct example",
			s: &struct {
				This int
				That string
				Bits []int
				Bobs []bool
			}{
				This: 345897,
				That: "testing123",
				Bits: []int{0, 55, 2, 1111, 145},
				Bobs: []bool{true, true, false},
			},
			expectedValues: []interface{}{
				345897,
				"testing123",
				[]int{0, 55, 2, 1111, 145},
				[]bool{true, true, false},
			},
			expectError: false,
		},
		{
			name: "Chat struct example",
			s: &dbmodels.Chat{
				ChatId:       "AIRFGA4I7Y8FRGYRF",
				SenderId:     "EWRGO2IEGG8257G",
				ReceiverId:   "WGUOIEYRBGOEIYBG0547",
				CreationDate: TEST_CREATION_DATE_2,
			},
			expectedValues: []interface{}{
				"AIRFGA4I7Y8FRGYRF",
				"EWRGO2IEGG8257G",
				"WGUOIEYRBGOEIYBG0547",
				TEST_CREATION_DATE_2,
			},
			expectError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := helpers.StructFieldValues(tc.s)

			for i, v := range values {
				if !reflect.DeepEqual(v, tc.expectedValues[i]) {
					t.Errorf("returned value is not the same as expected value")
				}
			}

			if tc.expectError {
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

func TestStructFieldAddress(t *testing.T) {
	// Create a sample struct and a pointer to it
	sampleStruct := TestAddressStruct{"hello", 42, true}
	ptrToStruct := &sampleStruct

	// Call the function under test
	addresses, err := helpers.StructFieldAddress(ptrToStruct)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the returned addresses
	expectedAddresses := []interface{}{
		&ptrToStruct.Field1,
		&ptrToStruct.Field2,
		&ptrToStruct.Field3,
	}

	// Make sure the addresses match the expected values
	if !reflect.DeepEqual(addresses, expectedAddresses) {
		t.Errorf("Unexpected addresses. Got %v, want %v", addresses, expectedAddresses)
	}
}

func TestDecideStructType(t *testing.T) {
	testcases := []struct {
		name           string
		table          string
		expectedStruct interface{}
		expectError    bool
	}{
		{
			name:           "Posts table",
			table:          "Posts",
			expectedStruct: &dbmodels.Post{},
			expectError:    false,
		},
		{
			name:           "Chats Messages table",
			table:          "Chats_Messages",
			expectedStruct: &dbmodels.ChatMessage{},
			expectError:    false,
		},
		{
			name:           "Followers table",
			table:          "Followers",
			expectedStruct: &dbmodels.Follower{},
			expectError:    false,
		},
		{
			name:           "Unknown table",
			table:          "Admins",
			expectedStruct: nil,
			expectError:    true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			obj, err := helpers.DecideStructType(tc.table)

			if !reflect.DeepEqual(tc.expectedStruct, obj) {
				t.Errorf("not the expected struct")
			}

			if tc.expectError {
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
