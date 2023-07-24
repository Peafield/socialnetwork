package dbutils

import (
	"fmt"
	"strings"
)

func UpdateConditionConstructor(Conditions map[string]interface{}) string {
	var conditionStatement string = "WHERE "
	var temp []string
	// conditionValues := make([]interface{}, 0)

	for k, v := range Conditions {
		if fmt.Sprintf("%T", v) == "string" {
			temp = append(temp, fmt.Sprintf(`%s = "%v"`, k, v))
		} else {
			temp = append(temp, fmt.Sprintf("%s = %v", k, v))
		}

	}

	conditionStatement += strings.Join(temp, " AND ")
	return conditionStatement
}
