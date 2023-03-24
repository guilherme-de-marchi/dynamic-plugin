package dypl

import (
	"fmt"
	"reflect"
)

type Struct struct {
	Receiver reflect.Value
	fields   map[string]reflect.Type
	methods  map[string]Method
}

func NewStruct(recv reflect.Value) (*Struct, error) {
	if !isStruct(recv) {
		return nil, ErrIsNotStruct
	}
	fs, err := getStructFields(recv)
	if err != nil {
		return nil, err
	}
	ms, err := getStructMethods(recv)
	return &Struct{
		Receiver: recv,
		fields:   fs,
		methods:  ms,
	}, err
}

func (s *Struct) Call(name string, args ...any) ([]reflect.Value, error) {
	m, ok := s.methods[name]
	if !ok {
		return nil, ErrNotFound
	}
	p := append([]reflect.Value{s.Receiver}, AnyToValue(args)...)
	return m.call(p), nil
}

func (s *Struct) ListMethods() []string {
	return keys(s.methods)
}

func (s *Struct) ListFields() []string {
	return keys(s.fields)
}

func (s *Struct) GetMethod(name string) (Method, error) {
	m, ok := s.methods[name]
	if !ok {
		return Method{}, ErrNotFound
	}
	return m, nil
}

func (s *Struct) GetField(name string) (reflect.Value, error) {
	_, ok := s.fields[name]
	if !ok {
		return reflect.Value{}, ErrNotFound
	}
	return reflect.Indirect(s.Receiver).FieldByName(name), nil
}

func (s *Struct) String() string {
	return fmt.Sprintf("%v <methods:%v> <fields:%v>", s.Receiver.Type().String(), s.methods, s.fields)
}

type Method struct {
	method        reflect.Method
	ExpIn, ExpOut []reflect.Type
}

func (m Method) call(args []reflect.Value) []reflect.Value {
	return m.method.Func.Call(args)
}

func (m Method) String() string {
	return fmt.Sprintf("%v", m.method.Func.Type().String())
}

func getStructFields(s reflect.Value) (map[string]reflect.Type, error) {
	if !isStruct(s) {
		return nil, ErrIsNotStruct
	}
	fs := make(map[string]reflect.Type)
	xf := reflect.Indirect(s)
	for i := 0; i < xf.NumField(); i++ {
		f := xf.Type().Field(i)
		fs[f.Name] = f.Type
	}
	return fs, nil
}

func getStructMethods(s reflect.Value) (map[string]Method, error) {
	if !isStruct(s) {
		return nil, ErrIsNotStruct
	}
	x := s.Type()
	ms := make(map[string]Method)
	for i := 0; i < x.NumMethod(); i++ {
		m := x.Method(i)
		expIn, err := getFuncInputs(m.Func)
		if err != nil {
			return nil, err
		}
		expOut, err := getFuncOutputs(m.Func)
		if err != nil {
			return nil, err
		}
		ms[m.Name] = Method{
			method: m,
			ExpIn:  expIn,
			ExpOut: expOut,
		}
	}
	return ms, nil
}
