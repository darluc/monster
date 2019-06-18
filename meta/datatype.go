package meta

import (
	"reflect"
)

// DataType defines basic data types in the system
type DataType interface {
	// TypeCheck reports whether value can be convert into the required datatype
	TypeCheck(value interface{}) bool
	ReflectType() reflect.Type
}

type TypeChecker func(value interface{}) bool

type Collector interface {
	Values() []interface{}
	Add(val interface{})
	Remove(val interface{})
}

type CollectionType interface {
	Collector
	DataType
}

// CompositeDataType is a datatype defined by meta object
type CompositeDataType interface {
	Object
	DataType
}
