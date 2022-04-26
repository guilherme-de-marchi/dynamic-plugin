package dynamic_plugin

import (
	"reflect"
)

func keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, len(m))
	var i int
	for k := range m {
		ks[i] = k
		i++
	}
	return ks
}

func AnyToAny[T any](args ...T) []any {
	vls := make([]any, len(args))
	for i, v := range args {
		vls[i] = v
	}
	return vls
}

func anyToValue(args []any) []reflect.Value {
	values := make([]reflect.Value, len(args))
	for i, v := range args {
		values[i] = reflect.ValueOf(v).Index(0)
	}
	return values
}
