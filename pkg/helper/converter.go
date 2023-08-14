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
		getField := structDataType.Field(i)
		if InArray(removeColumn, getField.Name) < 0 {
			fieldValue := structDataValue.Field(i)
			fieldValueToString := fmt.Sprintf("%v", fieldValue)
			newInterfaceLayout[getField.Tag.Get("json")] = fieldValueToString
		}
	}

	return newInterfaceLayout
}

func StructToInterfaceArrObj(structWithData any, removeColumn []string) []interface{} {
	// for _, val = range structWithData {

	// }

	var transformToArray []interface{}
	return transformToArray
	// newInterfaceLayout := make(map[string]interface{})

	// structDataType := reflect.TypeOf(structWithData)
	// structDataValue := reflect.ValueOf(structWithData)

	// for i := 0; i < structDataType.NumField(); i++ {
	// 	getField := structDataType.Field(i)
	// 	if InArray(removeColumn, getField.Name) < 0 {
	// 		fieldValue := structDataValue.Field(i)
	// 		fieldValueToString := fmt.Sprintf("%v", fieldValue)
	// 		newInterfaceLayout[getField.Tag.Get("json")] = fieldValueToString
	// 	}
	// 	transformToArray = append(transformToArray, newInterfaceLayout)
	// }

	// return transformToArray
}


