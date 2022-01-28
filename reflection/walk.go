package reflection

import "reflect"

type walkerFuncT func(input string)

func reflectValue(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, walkerFunc walkerFuncT) {
	val := reflectValue(x)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), walkerFunc)
		}
	} else {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			switch field.Kind() {
			case reflect.String:
				walkerFunc(field.String())
			case reflect.Struct:
				walk(field.Interface(), walkerFunc)
			}
		}
	}
}
