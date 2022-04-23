package main

import (
	"errors"
	"fmt"
	"reflect"
)

// ########### v EXAMPLE v ########### //

type Block struct {
	Owner, Content string
	Next           *Block
}

type Blockchain struct {
	Root *Block
}

func (bc *Blockchain) ListBlocks() []*Block {
	bs := make([]*Block, 0)
	curr := bc.Root
	for {
		if curr != nil {
			bs = append(bs, curr)
			curr = curr.Next
			continue
		}
		break
	}
	return bs
}

func (bc *Blockchain) FindOwner(name string) *Block {
	curr := bc.Root
	for {
		if curr != nil {
			if curr.Owner == name {
				return curr
			}
			curr = curr.Next
			continue
		}
		break
	}
	return nil
}

// ########### ^ EXAMPLE ^ ########### //

type Method struct {
	method  reflect.Method
	In, Out []reflect.Type
}

func (m Method) Call(args ...any) []reflect.Value {
	return m.method.Func.Call(anyToValue(args...))
}

type Methods map[string]Method

func (ms Methods) Call(name string, args ...any) ([]reflect.Value, error) {
	m, ok := ms[name]
	if !ok {
		return nil, errors.New("not found")
	}
	return m.Call(args...), nil
}

type Struct struct {
	Fields  map[string]reflect.Type
	Methods Methods
}

func getFuncInputs(f reflect.Type) ([]reflect.Type, error) {
	in := make([]reflect.Type, 0)
	if f.Kind() != reflect.Func {
		return nil, errors.New("received value is not a function")
	}
	for i := 0; i < f.NumIn(); i++ {
		x := f.In(i)
		in = append(in, x)
	}
	return in, nil
}

func getFuncOutputs(f reflect.Type) ([]reflect.Type, error) {
	out := make([]reflect.Type, 0)
	if f.Kind() != reflect.Func {
		return nil, errors.New("received value is not a function")
	}
	for i := 0; i < f.NumOut(); i++ {
		x := f.Out(i)
		out = append(out, x)
	}
	return out, nil
}

func getStructFields(s any) map[string]reflect.Type {
	fs := make(map[string]reflect.Type)
	xf := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < xf.NumField(); i++ {
		f := xf.Type().Field(i)
		fs[f.Name] = f.Type
	}
	return fs
}

func getStructMethods(s any) (Methods, error) {
	x := reflect.TypeOf(s)
	if reflect.Indirect(reflect.ValueOf(s)).Kind() != reflect.Struct && x.Kind() != reflect.Struct {
		return nil, errors.New("received value is not a struct")
	}
	ms := make(Methods)
	for i := 0; i < x.NumMethod(); i++ {
		m := x.Method(i)
		in, err := getFuncInputs(m.Func.Type())
		if err != nil {
			return nil, err
		}
		out, err := getFuncOutputs(m.Func.Type())
		if err != nil {
			return nil, err
		}
		ms[m.Name] = Method{
			method: m,
			In:     in,
			Out:    out,
		}
	}
	return ms, nil
}

func MapStruct(s any) (Struct, error) {
	ms, err := getStructMethods(s)
	return Struct{Fields: getStructFields(s), Methods: ms}, err
}

func main() {
	b4 := &Block{Owner: "rafaela", Content: "block 4"}
	b3 := &Block{Owner: "antonio", Content: "block 3", Next: b4}
	b2 := &Block{Owner: "maria", Content: "block 2", Next: b3}
	b1 := &Block{Owner: "carlos", Content: "block 1", Next: b2}
	bc := &Blockchain{Root: b1}
	s, err := MapStruct(bc)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(keys(s.Fields), keys(s.Methods))
	values, err := s.Methods.Call("FindOwner", bc, "carlos")
	if err != nil {
		panic(err)
	}
	for _, v := range values {
		fmt.Println(reflect.Indirect(v).FieldByName("Owner"))
	}
}
