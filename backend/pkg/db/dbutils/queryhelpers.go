package dbutils

import (
	"fmt"
	"strings"
)

/*
ConditionStatementConstructor constructs a condition statement for SQLite queries where multiple
statements should be satisfied.

All condition statements start with "WHERE ", after that, any of the conditions that are passed through
are formatted for SQLite and appended to a slice.  Then that slice is joined together with the " AND "
SQLite operator inbetween.

Parameters:
  - conditions (map[string]interface{}): the conditions for the query in map format.

Returns:
  - string: a condition statement string.
*/
func ConditionStatementConstructor(conditions map[string]interface{}) string {
	if len(conditions) == 0 {
		return ""
	}
	var conditionStatement string = "WHERE "
	var temp []string

	for k, v := range conditions {
		if fmt.Sprintf("%T", v) == "string" {
			temp = append(temp, fmt.Sprintf(`%s = "%v"`, k, v))
		} else {
			temp = append(temp, fmt.Sprintf(`%s = %v`, k, v))
		}
	}

	conditionStatement += strings.Join(temp, " AND ")
	return conditionStatement
}

/**/
func UpdateSetConstructor(MutableValues map[string]interface{}) string {
	if len(MutableValues) == 0 {
		return ""
	}

	var setStatement string = "SET "
	var Temp []string

	for k, v := range MutableValues {
		Temp = append(Temp, fmt.Sprintf("%s = '%v'", k, v))

	}
	setStatement += strings.Join(Temp, ", ")

	return setStatement
}
