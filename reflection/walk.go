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
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), walkerFunc)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), walkerFunc)
		}
	case reflect.String:
		walkerFunc(val.String())
	}
}
