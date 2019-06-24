package composite

import (
	"monster/meta"
	"monster/system/implement/base"
	"reflect"
)

// GenesisType is datatype of all instances
var GenesisType meta.CompositeDataType

func init() {
	genesisModel := base.NewBaseObject("__genesis__")
	GenesisType = &tCompositeType{typeModel: genesisModel, Instance: base.NewBaseInstance(genesisModel), super: nil}
}

type tCompositeType struct {
	meta.Instance
	typeModel meta.Object
	super     meta.CompositeDataType
}

func (t *tCompositeType) TypeCheck(value interface{}) (isValid bool) {
	switch value.(type) {
	case meta.Instance:
		ins := value.(meta.Instance)
		for _, fld := range t.typeModel.Fields() {
			if isValid = ins.MetaObject().HasField(fld); !isValid {
				return
			}
		}
	case meta.InstanceCollector:
		insCollection := value.(meta.InstanceCollector)
		for _, ins := range insCollection.Values() {
			for _, fld := range t.typeModel.Fields() {
				if isValid = ins.MetaObject().HasField(fld); !isValid {
					return
				}
			}
		}
	}
	return
}

func (*tCompositeType) ReflectType() reflect.Type {
	panic("implement me")
}

func (t *tCompositeType) TypeModel() meta.Object {
	return t.typeModel
}

// NewMetaDrivenType creates meta-driven datatype with a type model and a parent datatype
func NewMetaDrivenType(typeModel meta.Object, superType meta.CompositeDataType) meta.CompositeDataType {
	return &tCompositeType{Instance: base.NewBaseInstance(superType.TypeModel()), typeModel: typeModel, super: superType}
}

// NewDataType always creates 1st generation datatype as child of GenesisType
func NewDataType(typeModel meta.Object) meta.CompositeDataType {
	// @todo[medium] should check typeModel collision
	return NewMetaDrivenType(typeModel, GenesisType)
}

// ExtendsMetaDrivenType extends parent datatype's model with its own type model, and creates datatype from the newly
// created type model.
func ExtendsMetaDrivenType(typeModel meta.Object, superType meta.CompositeDataType) meta.CompositeDataType {
	extendedTypeModel := base.NewBaseObject("__" + typeModel.Name() + "__:" + superType.MetaObject().Name())
	meta.ExtendsObject(extendedTypeModel, typeModel, superType.TypeModel())
	return NewMetaDrivenType(extendedTypeModel, superType)
}
