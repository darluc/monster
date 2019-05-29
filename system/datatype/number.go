package datatype

import "reflect"

var (
	NumberType  tNumber
	numberKinds []reflect.Kind = []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64}
)

type tNumber struct {
}

func (n tNumber) TypeCheck(value interface{}) bool {
	kind := reflect.TypeOf(value).Kind()
	for _, k := range numberKinds {
		if k == kind {
			return true
		}
	}
	return false
}

func (n tNumber) ReflectType() reflect.Type {
	return reflect.TypeOf(NumberType)
}
