package dbutils

import (
	"fmt"
	"strings"
)

/**/
func ConditionStatementConstructor(Conditions map[string]interface{}) string {
	if len(Conditions) == 0 {
		return ""
	}
	var conditionStatement string = "WHERE "
	var temp []string
	// conditionValues := make([]interface{}, 0)

	for k, v := range Conditions {
		temp = append(temp, fmt.Sprintf(`%s = %v`, k, v))
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
