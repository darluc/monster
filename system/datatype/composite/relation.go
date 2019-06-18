package composite

import (
	"monster/meta"
	"monster/system/implement/base"
)

func NewRelationType(sourceObj meta.Object, targetObj meta.Object) *MetaDrivenType {
	relationObject := base.NewBaseObject()
	sourceField := base.NewBaseField("source", NewMetaDrivenType(sourceObj))
	targetField := base.NewBaseField("target", NewMetaDrivenType(targetObj))
	relationObject.AddField(sourceField)
	relationObject.AddField(targetField)
	return NewMetaDrivenType(relationObject)
}

func NewRelation(source meta.Instance, target meta.Instance, relation *MetaDrivenType) meta.Instance {
	ins := base.NewBaseInstance(relation)
	ins.SetFieldValue(relation.Field("source"), source)
	ins.SetFieldValue(relation.Field("target"), target)
	return ins
}
