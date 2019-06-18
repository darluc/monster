package composite

import (
	"monster/meta"
	"reflect"
)

// MetaDrivenType is a kind of datatype which are defined by meta object, and all their values are corresponding instances.
type MetaDrivenType struct {
	meta.Object
}

func (mt *MetaDrivenType) TypeCheck(value interface{}) bool {
	inputInstance, ok := value.(meta.Instance)
	if ok {
		// input instance's meta object is the same with the one which driven this datatype
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
