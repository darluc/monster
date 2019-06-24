package composite

import (
	"monster/meta"
	"reflect"
)

var (
	ObjectType, FieldType meta.DataType
)

func init() {
	ObjectType = &tObject{}
	FieldType = &tField{}
}

// tObject is type for meta object
type tObject struct {
	meta.Instance
}

func (*tObject) TypeCheck(value interface{}) bool {
	_, ok := value.(meta.Object)
	return ok
}

func (*tObject) ReflectType() reflect.Type {
	panic("implement me")
}

// tField is type for meta field
type tField struct {
	meta.Instance
}

func (*tField) TypeCheck(value interface{}) bool {
	_, ok := value.(meta.Field)
	return ok
}

func (*tField) ReflectType() reflect.Type {
	panic("implement me")
}
