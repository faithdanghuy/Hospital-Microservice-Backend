package gorm

import "reflect"

func OmitFields(
	data interface{},
	callBack func(fieldValue reflect.Value) bool,
) []string {
	var omitFields []string
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i)

		if callBack(fieldValue) {
			omitFields = append(omitFields, fieldName)
		}
	}

	return omitFields
}
