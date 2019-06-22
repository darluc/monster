package composite

import (
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
	"reflect"
)

var innerAccess meta.Property

func init() {
	innerAccess = base.NewBaseProperty(property.FieldAccessibility, property.InnerAccess)
}

type tSuperType struct {
	meta.Instance
}

func (t *tSuperType) TypeCheck(value interface{}) bool {
	// at first value must be a meta instance
	if ins, ok := value.(meta.Instance); ok {
		// value's meta object must have all the meta fields of datatype's definition
		for _, fld := range t.instanceDefinition().Fields() {
			if fld.PropertyValue(property.FieldAccessibility) == property.InnerAccess {
				continue
			}
			if !ins.MetaObject().HasField(fld) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (*tSuperType) ReflectType() reflect.Type {
	panic("implement me")
}

func (t *tSuperType) instanceDefinition() meta.Object {
	if obj := t.FieldValue(InstanceSchema); obj != nil {
		return obj.(meta.Object)
	} else {
		return GenesisObject
	}
}

// NewSuperType defined a data type based on specified meta object
func NewSuperType(definition meta.Object) meta.CompositeDataType {
	typeInstance := base.NewBaseInstance(definition)
	return &tSuperType{Instance: typeInstance}
}
