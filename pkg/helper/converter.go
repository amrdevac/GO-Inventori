package helper

import (
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
			// fieldValueToString := fmt.Sprintf("%v", fieldValue)
			// newInterfaceLayout[getField.Tag.Get("json")] = fieldValueToString
			newInterfaceLayout[getField.Tag.Get("json")] = fieldValue.Interface()
		}
	}

	return newInterfaceLayout
}

func ArrStructToInterfaceObj(arrStructWithData interface{}, removeColumn []string, isUseNameField bool) []map[string]interface{} {
	val := reflect.ValueOf(arrStructWithData)
	interfaceArray := make([]interface{}, val.Len())

	for index := range RangeArray(val.Len()) {
		interfaceArray[index] = val.Index(index).Interface()
	}

	newInterfaceLayout := make(map[string]interface{})
	var arrNewInterfaceLayout []map[string]interface{}

	for _, structWithData := range interfaceArray {

		// Get the datatype of struct
		structDataType := reflect.TypeOf(structWithData)
		// Get the datatype of struct
		structDataValue := reflect.ValueOf(structWithData)

		// Loop by counting all column
		for i := 0; i < structDataType.NumField(); i++ {
			getFieldName := structDataType.Field(i).Name
			dataFieldTag := structDataType.Field(i).Tag
			if InArray(removeColumn, getFieldName) < 0 {
				fieldValue := structDataValue.Field(i)
				if isUseNameField {
					newInterfaceLayout[getFieldName] = fieldValue.Interface()
				} else {
					newInterfaceLayout[dataFieldTag.Get("json")] = fieldValue.Interface()
				}
			}
		}
		arrNewInterfaceLayout = append(arrNewInterfaceLayout, newInterfaceLayout)

	}

	return arrNewInterfaceLayout
}
