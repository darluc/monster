package composite

import (
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	RelationSource = "source"
	RelationTarget = "target"
)

var relationTypes map[string]*MetaDrivenType

func init() {
	relationTypes = make(map[string]*MetaDrivenType)
}

func existsType(source meta.Object, target meta.Object, fieldName string) (*MetaDrivenType, string) {
	typeId := source.Name() + "-" + fieldName + "->" + target.Name() // @todo:low not good
	t, ok := relationTypes[typeId]
	if ok {
		return t, typeId
	}
	return nil, typeId
}

// NewRelationType creates a new relation type for every source & field & target combination
func NewRelationType(sourceObj meta.Object, targetObj meta.Object, fieldName string) (retType *MetaDrivenType) {
	var typeId string
	retType, typeId = existsType(sourceObj, targetObj, fieldName) // try to find out an existed relation type
	if retType == nil {
		relationObject := base.NewBaseObject()
		sourceField := base.NewBaseField(RelationSource, NewMetaDrivenType(sourceObj))
		throughProperty := base.NewBaseProperty(property.RelationFieldName, fieldName)
		sourceField.AddProperty(throughProperty)
		targetField := base.NewBaseField(RelationTarget, NewMetaDrivenType(targetObj))
		relationObject.AddField(sourceField)
		relationObject.AddField(targetField)
		retType = NewMetaDrivenType(relationObject)
		relationTypes[typeId] = retType // remember the relation type
	}
	return
}

func BuildRelationship(source meta.Instance, target meta.Instance, relation *MetaDrivenType) {
	ins := base.NewBaseInstance(relation)
	sourceField := relation.Field(RelationSource)
	ins.SetFieldValue(sourceField, source)
	ins.SetFieldValue(relation.Field(RelationTarget), target)

	fieldName := sourceField.PropertyValue(property.RelationFieldName)
	source.SetFieldValue(source.MetaObject().Field(fieldName.(string)), ins)
}
