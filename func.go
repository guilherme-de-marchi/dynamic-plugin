package dynamic_plugin

import (
	"reflect"
)

type Func struct {
	fun           reflect.Value
	ExpIn, ExpOut []reflect.Type
}

func NewFunc(f reflect.Value) (*Func, error) {
	if !isFunc(f) {
		return nil, ErrIsNotFunc
	}
	expIn, err := getFuncInputs(f)
	if err != nil {
		return nil, err
	}
	expOut, err := getFuncOutputs(f)
	if err != nil {
		return nil, err
	}

	return &Func{
		fun:    f,
		ExpIn:  expIn,
		ExpOut: expOut,
	}, nil
}

func (f Func) Call(args ...reflect.Value) []reflect.Value {
	return f.fun.Call(args)
}

func getFuncInputs(f reflect.Value) ([]reflect.Type, error) {
	if !isFunc(f) {
		return nil, ErrIsNotFunc
	}
	in := make([]reflect.Type, 0)
	for i := 0; i < f.Type().NumIn(); i++ {
		x := f.Type().In(i)
		in = append(in, x)
	}
	return in, nil
}

func getFuncOutputs(f reflect.Value) ([]reflect.Type, error) {
	if !isFunc(f) {
		return nil, ErrIsNotFunc
	}
	out := make([]reflect.Type, 0)
	for i := 0; i < f.Type().NumOut(); i++ {
		x := f.Type().Out(i)
		out = append(out, x)
	}
	return out, nil
}
