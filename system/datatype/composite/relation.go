package composite

import (
	"fmt"
	"monster/meta"
	"monster/system/datatype"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	labelRelationSource = "source"
	labelRelationTarget = "target"
	labelRelationType   = "type"
)

var RelationType meta.CompositeDataType
var relationTypes map[string]meta.CompositeDataType
var (
	fldCardinality, fldRelationJoint meta.Field
)
var relationTypeModel meta.Object

func init() {
	relationTypeModel = base.NewBaseObject("__relation__")
	fldCardinality = base.NewBaseField("cardinality", datatype.NumberType)
	relationTypeModel.AddField(fldCardinality)
	fldRelationJoint = base.NewBaseField("connector", FieldType)
	relationTypeModel.AddField(fldRelationJoint)
	RelationType = NewDataType(relationTypeModel)
	relationTypes = make(map[string]meta.CompositeDataType)
}

func relationIdName(sourceObj meta.Object, targetObj meta.Object, fieldName string,
	cardinality property.TRelationCardinality) string {
	//@todo[low] may not quite right
	return fmt.Sprintf("%s-%s->%s@%d", sourceObj.Name(), fieldName, targetObj.Name(), cardinality)
}

func MakeRelation(sourceObj meta.Object, targetObj meta.Object, fieldName string,
	cardinality property.TRelationCardinality) meta.CompositeDataType {
	relationName := relationIdName(sourceObj, targetObj, fieldName, cardinality)
	if r, ok := relationTypes[relationName]; !ok {
		// create data instance model
		relationModel := base.NewBaseObject(relationName)
		sourceField := base.NewBaseField(labelRelationSource, NewDataType(sourceObj))
		targetField := base.NewBaseField(labelRelationTarget, NewDataType(targetObj))
		typeField := base.NewBaseField(labelRelationType, RelationType)
		relationModel.AddField(sourceField)
		relationModel.AddField(targetField)
		relationModel.AddField(typeField)
		// create new relation type, and set relation's field value: cardinality & connector
		relationType := ExtendsMetaDrivenType(relationModel, RelationType)
		relationType.SetFieldValue(fldCardinality, cardinality)
		relationTypes[relationName] = relationType
		// create joint field
		jointField := base.NewBaseField(fieldName, relationType)
		sourceObj.AddField(jointField)
		relationType.SetFieldValue(fldRelationJoint, jointField)
		return relationType
	} else {
		if sourceObj.Field(fieldName) == nil { // never gets here
			sourceObj.AddField(r.FieldValue(fldRelationJoint).(meta.Field))
		}
		return r
	}
}

//region --- Relation datatype helper functions ---

// RelationTypeCheck checks if the composite datatype is a relation type
func RelationTypeCheck(relationTypeInstance meta.CompositeDataType) bool {
	return relationTypeInstance.MetaObject() == relationTypeModel
}

// RelationSourceField returns source field of the relation datatype
func RelationSourceField(relationTypeInstance meta.CompositeDataType) meta.Field {
	return relationTypeInstance.TypeModel().Field(labelRelationSource)
}

// RelationTargetField returns target field of the relation datatype
func RelationTargetField(relationTypeInstance meta.CompositeDataType) meta.Field {
	return relationTypeInstance.TypeModel().Field(labelRelationTarget)
}

// RelationSourceTypeModel returns source field type model of the relation datatype
func RelationSourceTypeModel(relationType meta.CompositeDataType) meta.Object {
	return relationType.TypeModel().Field(labelRelationSource).Type().(meta.CompositeDataType).TypeModel()
}

// RelationSourceTypeModel returns target field type model of the relation datatype
func RelationTargetTypeModel(relationType meta.CompositeDataType) meta.Object {
	return relationType.TypeModel().Field(labelRelationTarget).Type().(meta.CompositeDataType).TypeModel()
}

// RelationCardinality returns cardinality of the relation type
func RelationCardinality(relationType meta.CompositeDataType) property.TRelationCardinality {
	return relationType.FieldValue(fldCardinality).(property.TRelationCardinality)
}

// RelationCardinality returns joint field of the relation type
func RelationJointField(relationType meta.CompositeDataType) meta.Field {
	return relationType.FieldValue(fldRelationJoint).(meta.Field)
}

//endregion

//region --- Relation instance helper functions ---
func BuildRelationship(relationType meta.CompositeDataType, source meta.Instance, target meta.Instance) meta.Instance {
	relation := base.NewBaseInstance(relationType.TypeModel())
	relation.SetFieldValue(relationType.TypeModel().Field(labelRelationSource), source)
	relation.SetFieldValue(relationType.TypeModel().Field(labelRelationTarget), target)
	relation.SetFieldValue(relationType.TypeModel().Field(labelRelationType), relationType)
	return relation
}

//endregion
