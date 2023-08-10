package helper

import (
	"fmt"
	"reflect"
)

func StructToInterfaceObj(structWithData interface{}, removeColumn []string) map[string]interface{} {
	newInterfaceLayout := make(map[string]interface{})

	structDataType := reflect.TypeOf(structWithData)
	structDataValue := reflect.ValueOf(structWithData)

	for i := 0; i < structDataType.NumField(); i++ {
		getFieldName := structDataType.Field(i).Name
		dataTag := structDataType.Field(i).Tag
		if InArray(removeColumn, getFieldName) < 0 {
			fieldValue := structDataValue.Field(i)
			fieldValueToString := fmt.Sprintf("%v", fieldValue)
			newInterfaceLayout[dataTag.Get("json")] = fieldValueToString
		}
	}
	
	return newInterfaceLayout
}
