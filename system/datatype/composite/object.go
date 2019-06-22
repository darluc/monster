package composite

import (
	"monster/meta"
	"reflect"
)

var (
	ObjectType = tObject{}
)

// tObject is type of all composite data types
type tObject struct {
	meta.Instance
}

func (tObject) Definition() meta.Object {
	return meta.Object(nil)
}

// TypeCheck must be a meta
func (tObject) TypeCheck(value interface{}) bool {
	_, ok := value.(meta.Object)
	return ok
}

func (tObject) ReflectType() reflect.Type {
	return reflect.TypeOf(tObject{})
}
