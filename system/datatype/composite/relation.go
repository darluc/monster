package composite

import (
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	RelationSource = "source"
	RelationTarget = "target"
	RelationType   = "type"
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
func NewRelationType(sourceObj meta.Object, targetObj meta.Object, fieldName string,
	cardinality property.TRelationCardinality) (retType *MetaDrivenType) {
	var typeId string
	retType, typeId = existsType(sourceObj, targetObj, fieldName) // search for an existed relation type
	if retType == nil {
		relationObject := base.NewBaseObject("[relation]: " + typeId)
		retType = NewMetaDrivenType(relationObject)
		sourceField := base.NewBaseField(RelationSource, NewMetaDrivenType(sourceObj))
		targetField := base.NewBaseField(RelationTarget, NewMetaDrivenType(targetObj))
		typeField := base.NewBaseField(RelationType, GenesisType)
		relationObject.AddField(sourceField)
		relationObject.AddField(targetField)
		relationObject.AddField(typeField)
		// set fieldName as relation indicating name
		relationObject.AddProperty(base.NewBaseProperty(property.RelationIndicatingName, fieldName))
		// set relation cardinality property
		relationObject.AddProperty(base.NewBaseProperty(property.RelationCardinality, cardinality))
		relationTypes[typeId] = retType // remember the relation type
	}
	return
}

// BuildRelationship builds single relation instance between two instances with specified relation type
func BuildRelationship(source meta.Instance, target meta.Instance, relation *MetaDrivenType) meta.Instance {
	ins := base.NewBaseInstance(relation)
	ins.SetFieldValue(relation.Field(RelationSource), source)
	ins.SetFieldValue(relation.Field(RelationTarget), target)
	ins.SetFieldValue(relation.Field(RelationType), relation)
	return ins
}
