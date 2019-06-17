package executor

import (
	"context"
	"math/rand"
	"monster/meta"
	"monster/system/datatype"
	"monster/system/implement/base"
	"testing"
	"time"
)

func TestCreateExec(t *testing.T) {
	f1 := base.NewBaseField("name", datatype.StringType)
	f2 := base.NewBaseField("code", datatype.StringType)
	f3 := base.NewBaseField("count", datatype.NumberType)
	obj := base.NewBaseObject()
	obj.AddField(f1)
	obj.AddField(f2)
	obj.AddField(f3)

	exec := CreateExec{InstanceCreator: base.NewBaseInstance, MetaObject: obj}
	var err error
	err = exec.Open()
	instances := meta.NewBatch()
	instanceCount := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10) + 1
	if err == nil {
		for i := 0; i < instanceCount; i++ {
			err = exec.Next(context.Background(), instances)
		}
	}
	if err != nil {
		err = exec.Close()
	}
	if err != nil {
		t.Errorf("error occured: %v", err)
	} else if len(instances.Items) == instanceCount {
		t.Logf("%d instance created: %v", instanceCount, instances.Items)
	} else {
		t.Errorf("we got instances %v, count is different with %d", instances.Items, instanceCount)
	}
}
