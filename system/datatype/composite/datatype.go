package composite

import (
	"monster/meta"
	"reflect"
)

var (
	GenesisType *MetaDrivenType
)

func init() {
	GenesisType = NewMetaDrivenType(*new(meta.Object))
}

// MetaDrivenType is a kind of datatype which are defined by meta object, and all their values are corresponding instances.
type MetaDrivenType struct {
	meta.Object
}

func (mt *MetaDrivenType) TypeCheck(value interface{}) bool {
	if mt == GenesisType { // all meta object is type of Genesis
		_, ok := value.(meta.Object)
		return ok
	}

	switch value.(type) {
	case meta.Instance:
		inputInstance := value.(meta.Instance)
		if inputInstance.MetaObject() == mt.Object {
			return true
		}
		if insType, ok := inputInstance.MetaObject().(*MetaDrivenType); ok {
			return insType == mt
		}
	case meta.InstanceCollector:
		instanceCollection := value.(meta.InstanceCollector)
		if instanceCollection.Size() == 0 {
			return true
		}
		ins := instanceCollection.Values()[0]
		if ins.MetaObject() == mt.Object {
			return true
		}
		if insType, ok := ins.MetaObject().(*MetaDrivenType); ok {
			return insType == mt
		}
	case *MetaDrivenType:
		return value.(*MetaDrivenType).Object == mt.Object
	}
	return false
}

func (*MetaDrivenType) ReflectType() reflect.Type {
	return nil //@todo implement reflect.Type interface with meta object wrapper/converter
}

func NewMetaDrivenType(object meta.Object) *MetaDrivenType {
	return &MetaDrivenType{Object: object}
}
