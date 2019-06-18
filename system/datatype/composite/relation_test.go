package composite

import (
	"monster/system/datatype"
	"monster/system/implement/base"
	"testing"
)

func TestRelationType(t *testing.T) {
	objA := base.NewBaseObject()
	objB := base.NewBaseObject()

	relationFieldName := "parent"
	relationType := NewRelationType(objA, objB, relationFieldName)
	relationField := base.NewBaseField(relationFieldName, relationType)
	objA.AddField(relationField)

	nameField := base.NewBaseField("name", datatype.StringType)
	objA.AddField(nameField)
	objB.AddField(nameField)

	insA := base.NewBaseInstance(objA)
	insA.SetFieldValue(nameField, "George")
	insB := base.NewBaseInstance(objB)
	insB.SetFieldValue(nameField, "Daddy Piggy")
	BuildRelationship(insA, insB, relationType)

	t.Logf("%s's %s is %s", insA.FieldValue(nameField), relationField.Name(), insB.FieldValue(nameField))
}
