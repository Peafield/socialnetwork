package helpers

import (
	"reflect"
	"socialnetwork/pkg/models/dbmodels"
)

func StructFieldValues(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		if fieldType.Name != "CreationDate" {
			values = append(values, v.Field(i).Interface())
		}

	}
	return values
}

func StructFieldAddress(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	v = v.Elem() // de-reference the pointer to get the underlying struct
	var addresses []interface{}
	for i := 0; i < v.NumField(); i++ {
		addresses = append(addresses, v.Field(i).Addr().Interface())
	}
	return addresses
}

func DecideStructType(table string) interface{} {
	switch table {
	case "Users":
		var obj dbmodels.User
		return &obj
	}
	return nil
}
