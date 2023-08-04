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
SQLite operator inbetween.  The values are also appended and returned at the same time so the order is kept (Go
maps are random).

Parameters:
  - conditions (map[string]interface{}): the conditions for the query in map format.

Returns:
  - string: a condition statement string.
  - []interface{}: a slice of interface so that the values remain in order when constructing statements
*/
func ConditionStatementConstructor(conditions map[string]interface{}) (string, []interface{}) {
	if len(conditions) == 0 {
		return "", nil
	}
	var conditionStatement string = "WHERE "
	var temp []string
	var ConditionValues []interface{}

	for k, v := range conditions {

		temp = append(temp, fmt.Sprintf(`%s = ?`, k))
		ConditionValues = append(ConditionValues, v)
	}

	conditionStatement += strings.Join(temp, " AND ")
	return conditionStatement, ConditionValues
}

/**/
func UpdateSetConstructor(MutableValues map[string]interface{}) (string, []interface{}) {
	//if nothing received, return nothing
	if len(MutableValues) == 0 {
		return "", nil
	}

	var setStatement string = "SET "
	var Temp []string
	var ColumnValues []interface{}

	for k, v := range MutableValues {
		Temp = append(Temp, fmt.Sprintf("%s = ?", k))
		ColumnValues = append(ColumnValues, v)

	}
	setStatement += strings.Join(Temp, ", ")

	return setStatement, ColumnValues
}
