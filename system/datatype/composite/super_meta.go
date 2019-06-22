package composite

import (
	"monster/meta"
	"monster/system/implement/base"
)

var (
	InstanceSchema meta.Field
)

// tSuperMetaType is the super type of all other meta driven data types
var tSuperMetaType meta.CompositeDataType

func init() {
	metaDrivenTypeDefinition := base.NewBaseObject("__meta_datatype__")
	// based meta object
	InstanceSchema = base.NewBaseField("__object__", ObjectType)
	metaDrivenTypeDefinition.AddField(InstanceSchema)
	tSuperMetaType = Extends(metaDrivenTypeDefinition, GenesisInstanceType)
}
