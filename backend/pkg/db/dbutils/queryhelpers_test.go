package dbutils_test

import (
	"reflect"
	"socialnetwork/pkg/db/dbutils"
	"strings"
	"testing"
)

func TestUpdateConditionStatement(t *testing.T) {
	testcases := []struct {
		CaseName       string
		Conditions     map[string]interface{}
		ExpectedOutput string
	}{
		{
			CaseName:       "provided only 1 condition",
			Conditions:     map[string]interface{}{"column1": "value1"},
			ExpectedOutput: "WHERE column1 = value1",
		},
		{
			CaseName:       "provided a condition with non-string value",
			Conditions:     map[string]interface{}{"column1": 123},
			ExpectedOutput: "WHERE column1 = 123",
		},
		{
			CaseName:       "provided more than 1 condition",
			Conditions:     map[string]interface{}{"column1": "value1", "column2": "value2", "column3": "value3"},
			ExpectedOutput: "WHERE column1 = value1 AND column2 = value2 AND column3 = value3",
		},
		{
			CaseName:       "provided 0 condition",
			Conditions:     map[string]interface{}{},
			ExpectedOutput: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.CaseName, func(t *testing.T) {
			result := dbutils.UpdateConditionConstructor(tc.Conditions)
			sep := "AND"

			//we split the expected output and the result by "AND"
			//then apply a comparison on the basis of reflect.DeepEqual
			//due to the fact that maps iteration occur randomly.
			if !reflect.DeepEqual(strings.Split(result, sep), strings.Split(tc.ExpectedOutput, sep)) {
				t.Error("Expected: ", tc.ExpectedOutput, " but got: ", result)
			}

		})
	}
}

func TestUpdateSetStatement(t *testing.T) {
	testcases := []struct {
		CaseName        string
		AffectedColumns map[string]interface{}
		ExpectedOutput  string
	}{
		{
			CaseName:        "provided only 1 column",
			AffectedColumns: map[string]interface{}{"column1": "value1"},
			ExpectedOutput:  "SET column1 = 'value1'",
		},
		{
			CaseName:        "provided a column with non-string value",
			AffectedColumns: map[string]interface{}{"column1": 123},
			ExpectedOutput:  "SET column1 = '123'",
		},
		{
			CaseName:        "provided more than 1 column",
			AffectedColumns: map[string]interface{}{"column1": "value1", "column2": 123, "column3": "value3"},
			ExpectedOutput:  "SET column1 = 'value1', column2 = '123', column3 = 'value3'",
		},
		{
			CaseName:        "provided 0 column",
			AffectedColumns: map[string]interface{}{},
			ExpectedOutput:  "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.CaseName, func(t *testing.T) {
			result := dbutils.UpdateSetConstructor(tc.AffectedColumns)
			sep := ", "

			//we split the expected output and the result by "AND"
			//then apply a comparison on the basis of reflect.DeepEqual
			//due to the fact that maps iteration occur randomly.
			if !reflect.DeepEqual(strings.Split(result, sep), strings.Split(tc.ExpectedOutput, sep)) {
				t.Error("Expected: ", tc.ExpectedOutput, " but got: ", result)
			}

		})
	}
}
