package datatype

import (
	"reflect"
)

var (
	StringType, NumberType reflect.Type
)

func init() {
	StringType = reflect.TypeOf("")
	NumberType = reflect.TypeOf(1)
}
