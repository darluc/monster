package datatype

import (
	"reflect"
)

var (
	StringType, NumberType reflect.Type
	// RelationType is actually the type of instance type, so we must setup the type while initializing the system
	RelationType reflect.Type
)

func init() {
	StringType = reflect.TypeOf("")
	NumberType = reflect.TypeOf(1)
}

func SetRelationDataType(t reflect.Type) {
	RelationType = t
}
