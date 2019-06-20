package relation

import (
	"context"
	"math/rand"
	"monster/meta"
	"monster/system"
	"monster/system/datatype"
	"monster/system/datatype/composite"
	"monster/system/implement/base"
	"monster/system/property"
	"strconv"
	"testing"
	"time"
)

func init() {
	system.Startup()
}

func TestBuildExec_Open(t *testing.T) {
	exec := BuildExec{}
	err := exec.Open()
	if err != nil {
		t.Log(err)
	} else {
		t.Error("error not detected")
	}
}

func prepareObjectsAndRelation() (sourceObj meta.Object, targetObj meta.Object, relationType *composite.MetaDrivenType) {
	sourceObj = base.NewBaseObject("source")
	targetObj = base.NewBaseObject("target")

	nameField := base.NewBaseField("name", datatype.StringType)
	relationType = composite.NewRelationType(sourceObj, targetObj, "parent", property.OneToOne)
	parentField := base.NewBaseField("parent", relationType)

	sourceObj.AddField(nameField)
	sourceObj.AddField(parentField)
	targetObj.AddField(nameField)

	return
}

func TestBuildExec_OneToOne(t *testing.T) {
	sourceObj, targetObj, relationType := prepareObjectsAndRelation()

	parent := base.NewBaseInstance(targetObj)
	parent.SetFieldValue(parent.MetaObject().Field("name"), "Steve")

	childCount := 10
	children := make([]meta.Instance, 0, childCount)
	for i := 0; i < childCount; i++ {
		child := base.NewBaseInstance(sourceObj)
		child.SetFieldValue(child.MetaObject().Field("name"), "Bob"+strconv.Itoa(i))
		children = append(children, child)
	}

	exec := BuildExec{Relation: relationType, Target: parent}
	exec.Open()
	exec.Next(context.Background(), &meta.Batch{Items: children})
	exec.Close()

	rand.Seed(time.Now().UnixNano())
	lastPick := -1
	for i := 0; i < 2; i++ { // random pick two children to test their parent's name
		randPick := rand.Intn(10)
		if lastPick == randPick {
			i = i + 1
			continue
		}
		lastPick = randPick

		childRelation, ok := children[randPick].FieldValue(sourceObj.Field("parent")).(meta.Instance)
		if ok {
			relationTarget := childRelation.FieldValue(childRelation.MetaObject().Field(composite.RelationTarget)).(meta.Instance)
			relationSource := childRelation.FieldValue(childRelation.MetaObject().Field(composite.RelationSource)).(meta.Instance)
			if relationTarget == parent && relationSource == children[randPick] {
				t.Logf("%s's %s is %s", relationSource.FieldValue(relationSource.MetaObject().Field("name")),
					relationType.PropertyValue(property.RelationIndicatingName),
					relationTarget.FieldValue(relationTarget.MetaObject().Field("name")))
				continue
			} else {
				t.Errorf("relation target object is different from %v", targetObj)
			}
		} else {
			t.Fail()
		}
	}
}
