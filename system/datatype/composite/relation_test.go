package composite

import (
	"monster/system/datatype"
	"monster/system/implement/base"
	"monster/system/property"
	"testing"
)

func TestBuildRelationship(t *testing.T) {
	objChild := base.NewBaseObject("child")
	nameField := base.NewBaseField("name", datatype.StringType)
	objChild.AddField(nameField)

	objParent := base.NewBaseObject("parent")
	objParent.AddField(nameField)
	childRelation := NewRelationType(objParent, objChild, "children", property.OneToMany)
	childrenField := base.NewBaseField("children", childRelation)
	objParent.AddField(childrenField)

	parent := base.NewBaseInstance(objParent)
	parent.SetFieldValue(nameField, "Jim")
	child := base.NewBaseInstance(objChild)
	child.SetFieldValue(nameField, "bob")

	relation := BuildRelationship(childRelation, parent, child)
	parent.SetFieldValue(childrenField, relation)

	if childrenField != RelationJointField(childRelation) {
		t.Errorf("joint field is different %v, %v", childrenField, RelationJointField(childRelation))
	}

	if sourceIns := relation.FieldValue(RelationSourceField(childRelation)); sourceIns != parent {
		t.Errorf("relationship's source instance should be %v, but got %v", parent, sourceIns)
	}

	if targetIns := relation.FieldValue(RelationTargetField(childRelation)); targetIns != child {
		t.Errorf("relationship's target instance should be %v, but got %v", child, targetIns)
	}
}
