package composite

import (
	"monster/meta"
	"reflect"
)

type tObjectType struct {
	object *meta.Object
}

func (tObjectType) TypeCheck(value interface{}) bool {
	_, ok := value.(meta.Object)
	return ok
}

func (tObjectType) ReflectType() reflect.Type {
	panic("implement me")
}

var (
	GenesisObjectType tObjectType
)

func init() {
	GenesisObjectType = tObjectType{new(meta.Object)}
}
