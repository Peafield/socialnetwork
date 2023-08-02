package dbutils_test

import (
	"reflect"
	"socialnetwork/pkg/db/dbutils"
	"strings"
	"testing"
)

func TestUpdateConditionStatement(t *testing.T) {
	testcases := []struct {
		CaseName                string
		Conditions              map[string]interface{}
		ExpectedConditionStmnt  string
		ExpectedConditionValues []interface{}
	}{
		{
			CaseName:                "provided only 1 condition",
			Conditions:              map[string]interface{}{"column1": "value1"},
			ExpectedConditionStmnt:  "WHERE column1 = ?",
			ExpectedConditionValues: []interface{}{"value1"},
		},
		{
			CaseName:                "provided a condition with a non-string value",
			Conditions:              map[string]interface{}{"column1": 123},
			ExpectedConditionStmnt:  "WHERE column1 = ?",
			ExpectedConditionValues: []interface{}{123},
		},
		{
			CaseName:                "provided more than 1 condition",
			Conditions:              map[string]interface{}{"column1": "value1", "column2": 123, "column3": "value3"},
			ExpectedConditionStmnt:  "WHERE column1 = ? AND column2 = ? AND column3 = ?",
			ExpectedConditionValues: []interface{}{"value1", 123, "value3"},
		},
		{
			CaseName:                "provided 0 condition",
			Conditions:              map[string]interface{}{},
			ExpectedConditionStmnt:  "",
			ExpectedConditionValues: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.CaseName, func(t *testing.T) {
			result, values := dbutils.ConditionStatementConstructor(tc.Conditions)
			sep := "AND"

			//we split the expected output and the result by "AND"
			//then apply a comparison on the basis of reflect.DeepEqual
			//due to the fact that maps iteration occur randomly.
			resultComparison := reflect.DeepEqual(strings.Split(result, sep), strings.Split(tc.ExpectedConditionStmnt, sep))
			valuesComparison := reflect.DeepEqual(values, tc.ExpectedConditionValues)

			if !resultComparison {
				t.Error("Expected: ", tc.ExpectedConditionStmnt, " but got: ", result)
			}

			if !valuesComparison {
				t.Error("Expected: ", tc.ExpectedConditionValues, " but got: ", values)
			}
		})
	}
}

func TestUpdateSetStatement(t *testing.T) {
	testcases := []struct {
		CaseName             string
		AffectedColumns      map[string]interface{}
		ExpectedSetStatement string
		ExpectedSetValues    []interface{}
	}{
		{
			CaseName:             "provided only 1 column",
			AffectedColumns:      map[string]interface{}{"column1": "value1"},
			ExpectedSetStatement: "SET column1 = ?",
			ExpectedSetValues:    []interface{}{"value1"},
		},
		{
			CaseName:             "provided a column with a non-string value",
			AffectedColumns:      map[string]interface{}{"column1": 123},
			ExpectedSetStatement: "SET column1 = ?",
			ExpectedSetValues:    []interface{}{123},
		},
		{
			CaseName:             "provided more than 1 column",
			AffectedColumns:      map[string]interface{}{"column1": "value1", "column2": 123, "column3": "value3"},
			ExpectedSetStatement: "SET column1 = ?, column2 = ?, column3 = ?",
			ExpectedSetValues:    []interface{}{"value1", 123, "value3"},
		},
		{
			CaseName:             "provided 0 column",
			AffectedColumns:      map[string]interface{}{},
			ExpectedSetStatement: "",
			ExpectedSetValues:    nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.CaseName, func(t *testing.T) {
			result, values := dbutils.UpdateSetConstructor(tc.AffectedColumns)
			sep := ", "

			//we split the expected output and the result by ", "
			//then apply a comparison on the basis of reflect.DeepEqual
			//due to the fact that maps iteration occur randomly.
			resultComparison := reflect.DeepEqual(strings.Split(result, sep), strings.Split(tc.ExpectedSetStatement, sep))
			valuesComparison := reflect.DeepEqual(values, tc.ExpectedSetValues)

			if !resultComparison {
				t.Error("Expected: ", tc.ExpectedSetStatement, " but got: ", result)
			}

			if !valuesComparison {
				t.Error("Expected: ", tc.ExpectedSetValues, " but got: ", values)
			}

		})
	}
}
