package composite

import (
	"monster/meta"
	"reflect"
)

type MetaDrivenType struct {
	meta.Object
}

// TypeCheck checks whether value is type of MetaDrivenType
func (mt *MetaDrivenType) TypeCheck(value interface{}) bool {
	inputInstance, ok := value.(meta.Instance)
	if ok {
		if inputInstance.MetaObject() == mt.Object {
			return true
		}
		if insType, ok := inputInstance.MetaObject().(*MetaDrivenType); ok {
			return insType == mt
		}
	}
	return false
}

func (*MetaDrivenType) ReflectType() reflect.Type {
	return nil //@todo implement reflect.Type interface with meta object wrapper/converter
}

func NewMetaDrivenType(object meta.Object) *MetaDrivenType {
	return &MetaDrivenType{Object: object}
}
