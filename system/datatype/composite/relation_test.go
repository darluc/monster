package composite

import (
	"monster/system/implement/base"
	"testing"
)

func TestRelationType(t *testing.T) {
	objA := base.NewBaseObject()
	objB := base.NewBaseObject()

	relationType := NewRelationType(objA, objB)
	relationField := base.NewBaseField("parent", relationType)
	objA.AddField(relationField)

	insA := base.NewBaseInstance(objA)
	insB := base.NewBaseInstance(objB)

	relation := NewRelation(insA, insB, relationType)
	insA.SetFieldValue(relationField, relation)

	t.Logf("insA value: %v", insA)
}
