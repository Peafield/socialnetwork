package helpers

import (
	"fmt"
	"reflect"
	"socialnetwork/pkg/models/dbmodels"
	"time"
)

/*
StructFieldValues takes a pointer to a struct and returns the values of all of its properties.

It uses the reflect package to find the field names and their values and returns the relevant ones.

Parameters:
- s (interface{}): should be a pointer to a struct.

Returns:
- []interface{}: returns an array of different values that was inside the given struct.

Errors:
- if the input if not a pointer to a struct.
*/
func StructFieldValues(s interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		if fieldType.Name != "CreationDate" {
			// Normalize time fields before appending to values
			if v.Field(i).Type() == reflect.TypeOf(time.Time{}) {
				values = append(values, NormalizeTime(v.Field(i).Interface().(time.Time)))
			} else {
				values = append(values, v.Field(i).Interface())
			}
		}

	}
	return values, nil
}

/*
StructFieldAddress takes a pointer to a struct and returns the addresses of all of its properties.

It uses the reflect package to find the field names and their addresses and returns them.

Parameters:
- s (interface{}): should be a pointer to a struct.

Returns:
- []interface{}: returns an array of different addresses for the properties of the given struct.

Errors:
- if the input if not a pointer to a struct.
*/
func StructFieldAddress(s interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var addresses []interface{}
	for i := 0; i < v.NumField(); i++ {
		addresses = append(addresses, v.Field(i).Addr().Interface())
	}
	return addresses, nil
}

/*
DecideStructType takes a table type as a string and returns the relevant model.

It uses a switch case to cycle through the different options, then creates an empty variable of the relevant struct.

Parameters:
- table (string): the table name.

Returns:
- interface{}: a struct of the relevant table
- error: if the table inputted yields no results

Errors:
- if the table doesn't exist or there is no corresponding struct
*/
func DecideStructType(table string) (interface{}, error) {
	switch table {
	case "Users":
		var obj dbmodels.User
		return &obj, nil
	case "Posts":
		var obj dbmodels.Post
		return &obj, nil
	case "Chats":
		var obj dbmodels.Chat
		return &obj, nil
	case "Followers":
		var obj dbmodels.Follower
		return &obj, nil
	case "Groups":
		var obj dbmodels.Group
		return &obj, nil
	case "Groups_Events":
		var obj dbmodels.GroupEvent
		return &obj, nil
	case "Groups_Events_Attendees":
		var obj dbmodels.GroupEventAttendee
		return &obj, nil
	case "Groups_Members":
		var obj dbmodels.GroupMember
		return &obj, nil
	case "Notifications":
		var obj dbmodels.Notification
		return &obj, nil
	case "Chats_Messages":
		var obj dbmodels.ChatMessage
		return &obj, nil
	case "Comments":
		var obj dbmodels.Comment
		return &obj, nil
	case "Reactions":
		var obj dbmodels.Reaction
		return &obj, nil
	case "Sessions":
		var obj dbmodels.Session
		return &obj, nil
	default:
		return nil, fmt.Errorf("no valid struct type for that table")
	}
}
