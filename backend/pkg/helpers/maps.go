package helpers

import (
	"reflect"
	"strings"
)

func MapKeyContains(m map[string]interface{}, values []string) bool {
	for k := range m {
		if strings.Contains(strings.Join(values, " "), k) {
			return true
		}
	}
	return false
}

/*
TableColumnNames returns a slice of string representing all of the columns from a given DB table.
this function can be used, despite a table undergoing DB migration.

Parameters:
- m map[string]interface{}: a set of key/value pairs representing data (SQL statements Conditions or Columns to be affected).

- values []string: a set of strings representing data

Returns:
- bool: return true if the keys and the slice of strings are deeply equal

Example:
- Compares a set of keys in a map to match the values in the slice of strings.
*/
func ValuesMapComparison(m map[string]interface{}, values []string) bool {
	temp := []string{}

	for k := range m {
		temp = append(temp, k)
	}

	return reflect.DeepEqual(values, temp)
}
