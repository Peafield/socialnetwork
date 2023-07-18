package helpers

import (
	"fmt"
	"reflect"
)

/*
IsValidStructField makes sure a field is defined within a certain struct.

Uses reflect to get all the field name values, then loops through them to see if a match is found.
*/
func IsValidStructField(field, structType interface{}) (bool, error) {
	values := reflect.ValueOf(structType)

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i) == field {
			return true, nil
		}
	}

	return false, fmt.Errorf("not a valid struct field")
}
