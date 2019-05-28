package meta

import "reflect"

// DataType defines basic data types in the system
type DataType interface {
	// TypeCheck reports whether value can be convert into the required datatype
	TypeCheck(value interface{}) bool
	ReflectType() reflect.Type
}
