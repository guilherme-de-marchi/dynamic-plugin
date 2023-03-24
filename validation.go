package dypl

import "reflect"

func isStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Struct || reflect.Indirect(v).Kind() == reflect.Struct
}

func isFunc(v reflect.Value) bool {
	return v.Kind() == reflect.Func
}
