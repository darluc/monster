package composite

import (
	"fmt"
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	LabelRelationSource = "source"
	LabelRelationTarget = "target"
	LabelRelationType   = "type"
)

var relationTypes map[string]*MetaDrivenType

func init() {
	relationTypes = make(map[string]*MetaDrivenType)
}

func existsType(source meta.Object, target meta.Object, fieldName string, cardinality property.TRelationCardinality) (*MetaDrivenType, string) {
	typeId := fmt.Sprintf("R:%s-%s->%s:%d", source.Name(), fieldName, target.Name(), cardinality) //@todo[low] may not right
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
	retType, typeId = existsType(sourceObj, targetObj, fieldName, cardinality) // search for an existed relation type
	if retType == nil {
		relationObject := base.NewBaseObject("[relation]: " + typeId)
		retType = NewMetaDrivenType(relationObject)
		sourceField := base.NewBaseField(LabelRelationSource, NewMetaDrivenType(sourceObj))
		targetField := base.NewBaseField(LabelRelationTarget, NewMetaDrivenType(targetObj))
		typeField := base.NewBaseField(LabelRelationType, NewMetaDrivenType(relationObject))
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
	ins.SetFieldValue(relation.Field(LabelRelationSource), source)
	ins.SetFieldValue(relation.Field(LabelRelationTarget), target)
	ins.SetFieldValue(relation.Field(LabelRelationType), relation)
	return ins
}
