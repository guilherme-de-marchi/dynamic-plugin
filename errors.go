package dynamic_plugin

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrIsNotFunc   = errors.New("value is not a function")
	ErrIsNotStruct = errors.New("value is not a struct")
)
