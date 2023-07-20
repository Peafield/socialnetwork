package userdb

// // should store the immutable columns in a map

// func UpdateUserInfo(db *sql.DB, tableName string, Conditions map[string]interface{}, MutableValues map[string]interface{}) error {
// 	//maybe check the fields first before accessing them

// 	//temporary placeholder for the conditions
// 	var Temp []string

// 	for k, v := range MutableValues {
// 		Temp = append(TempConditions, fmt.Sprintf("%s = %v", k, v))
// 	}

// 	var UpdatedValues string = strings.Join(Temp, ", ")

// 	Query := fmt.Sprintf(`UPDATE %s SET %s WHERE user_id = %s;`, tableName, UpdatedValues)
// 	statement, err := db.Prepare(Query)

// 	if err != nil {
// 		return fmt.Errorf("failed to prepare UpdateUserInfo statement: %w", err)
// 	}

// 	defer statement.Close()

// 	result, err := statement.Exec()

// 	if err != nil {
// 		return fmt.Errorf("failed to execute UpdateUserInfo statement: %w", err)
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("failed to retrieve affected rows in UpdateUserInfo: %s", err)
// 	}

// 	if rowsAffected == 0 {
// 		return fmt.Errorf("no rows affected in UpdateUserInfo: %s", err)
// 	}

// 	return nil
// }

// func GetKeyValue(data string) map[string]interface{} {
// 	var result map[string]interface{}

// 	return result
// }
