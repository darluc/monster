package composite

import (
	"fmt"
	"monster/meta"
	"monster/system/datatype"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	LabelRelationSource = "source"
	LabelRelationTarget = "target"
)

var relationTypes map[string]meta.CompositeDataType
var RelationCardinality meta.Field
var RelationFieldName meta.Field

func init() {
	RelationFieldName = base.NewBaseField("pipeline", datatype.StringType)
	RelationCardinality = base.NewBaseField("cardinality", datatype.NumberType)
}

func relationIdName(sourceObj meta.Object, targetObj meta.Object, fieldName string,
	cardinality property.TRelationCardinality) string {
	return fmt.Sprintf("%s-%s->%s@%d", sourceObj.Name(), fieldName, targetObj.Name(), cardinality) //@todo[low] may not right
}

// NewRelationType creates a new relation type for every source & field & target combination
// {source:xxxObj, target:xxxObj, _cardinality: 1to1, _pipeline: "fieldX"}
func NewRelationType(sourceObj meta.Object, targetObj meta.Object, fieldName string,
	cardinality property.TRelationCardinality) meta.CompositeDataType {
	relationName := relationIdName(sourceObj, targetObj, fieldName, cardinality)
	if r, ok := relationTypes[relationName]; !ok {
		relationDefinition := base.NewBaseObject(relationName)
		sourceField := base.NewBaseField(LabelRelationSource, NewMetaType(sourceObj))
		targetField := base.NewBaseField(LabelRelationTarget, NewMetaType(targetObj))
		relationDefinition.AddField(sourceField)
		relationDefinition.AddField(targetField)
		relationDefinition.AddField(RelationCardinality)
		relationDefinition.AddField(RelationFieldName)
		//@todo[low] need a field of coupling type ?

		relationType := Extends(relationDefinition, GenesisInstanceType)
		relationType.SetFieldValue(sourceField, sourceObj)
		relationType.SetFieldValue(targetField, targetObj)
		relationType.SetFieldValue(RelationCardinality, cardinality)
		relationType.SetFieldValue(RelationFieldName, fieldName)
		relationTypes[relationName] = relationType
		return relationType
	} else {
		return r
	}
}

func sourceField(relationType meta.CompositeDataType) meta.Field {
	return relationType.MetaObject().Field(LabelRelationSource)
}

func targetField(relationType meta.CompositeDataType) meta.Field {
	return relationType.MetaObject().Field(LabelRelationTarget)
}

// BuildRelationship builds single relation instance between two instances with specified relation type
func BuildRelationship(source meta.Instance, target meta.Instance, relation meta.CompositeDataType) meta.Instance {
	ins := base.NewBaseInstance(relation.Definition())
	ins.SetFieldValue(sourceField(relation), source)
	ins.SetFieldValue(targetField(relation), target)
	return ins
}
