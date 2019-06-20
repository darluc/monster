package composite

import (
	"monster/system/implement/base"
	"monster/system/property"
	"testing"
)

func TestNewRelationType(t *testing.T) {
	objA := base.NewBaseObject("A")
	objB := base.NewBaseObject("B")
	relationName := "testing"
	relationType1 := NewRelationType(objA, objB, relationName, property.OneToOne)
	relationType2 := NewRelationType(objA, objB, relationName, property.OneToOne)
	if relationType1 == relationType2 {
		t.Logf("got the same relation type for same object relationship: %v, %v", relationType1, relationType2)
	} else {
		t.FailNow()
	}

	relationType3 := NewRelationType(objA, objB, relationName, property.OneToMany)
	if relationType1 == relationType3 {
		t.Errorf("got the same relation type for different caidinality: %v, %v", relationType1, relationType3)
	} else {
		t.Logf("got different relation type for different cardinality: %v, %v", relationType1, relationType3)
	}
}
