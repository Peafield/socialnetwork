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

func UndesiredParam(m map[string]interface{}, expectedParameters []string) bool {
	for _, v := range expectedParameters {
		_, exists := m[v]
		if !exists {
			return true
		}
	}
	return false
}

/*
ValuesMapComparison returns whether the arguments received are deeply equal.

Parameters:
- m map[string]interface{}: a set of key/value pairs representing data (SQL statements Conditions or Columns to be affected).

- values []string: a set of strings representing data

Returns:
- bool: return true if the keys and the slice of strings are deeply equal
*/
func ValuesMapComparison(m map[string]interface{}, values []string) bool {
	temp := []string{}

	for k := range m {
		temp = append(temp, k)
	}

	return reflect.DeepEqual(values, temp)
}
