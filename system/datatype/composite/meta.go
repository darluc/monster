package composite

import "monster/meta"

// NewMetaType creates new datatype with a meta object
func NewMetaType(object meta.Object) meta.CompositeDataType {
	typeIns := Extends(object, tSuperMetaType)
	typeIns.SetFieldValue(InstanceSchema, object)
	return typeIns
}
