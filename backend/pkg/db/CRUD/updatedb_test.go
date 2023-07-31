package crud_test

import (
	"database/sql"
	"regexp"
	"testing"
)

// what do i want to test:
// 1. if the query matches the required syntax
// 2. (optional) if the query contains the correct columns associated with the table
// 3. if the errors function correctly
// 4. to see whether rows affected by the query match the expected
type MockDB struct{}

// mock db
// mock prepare query
// mock execute statement
// rows affected by the query to be seen tomorrow
// regex update query syntax:
var UpdateFragment = `UPDATE\s{1}`
var TableNameFragment = `(\w\s){1}`
var SetStatementFragment = `((\w\s)=(\w)) `
var QueryRegex = regexp.MustCompile(`^UPDATE\s{1}(\w){1}(\sSET\s)$`)

func TestUpdateRowInfo(t *testing.T) {
	testcases := []struct {
		CaseName    string
		TableName   string
		Database    *sql.DB
		ExpectedErr bool
	}{}

	for _, tc := range testcases {
		t.Run(tc.CaseName, func(t *testing.T) {

		})
	}
}
