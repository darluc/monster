package composite

import (
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
)

var (
	GenesisObject       meta.Object
	GenesisInstanceType meta.CompositeDataType
)

func init() {
	GenesisObject = base.NewBaseObject("__genesis__")
	GenesisInstanceType = NewSuperType(GenesisObject)
}

// Extends creates new datatype out of meta object, and extends base datatype
func Extends(subtypeDefinition meta.Object, superType meta.CompositeDataType) (newDataType meta.CompositeDataType) {
	typeDefinitionObject := base.NewBaseObject("__datatype__" + subtypeDefinition.Name() + ":" + superType.MetaObject().Name())
	for _, fld := range superType.MetaObject().Fields() {
		if fld.PropertyValue(property.FieldAccessibility) != property.InnerAccess {
			typeDefinitionObject.AddField(fld)
		}
	}
	for _, fld := range subtypeDefinition.Fields() {
		typeDefinitionObject.AddField(fld)
	}
	typeIns := NewSuperType(typeDefinitionObject)
	return typeIns
}
