package datatype

import "reflect"

var (
	StringType tString
)

// tString is type of string
type tString struct {
}

func (s tString) TypeCheck(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}

func (s tString) ReflectType() reflect.Type {
	return reflect.TypeOf(StringType)
}
