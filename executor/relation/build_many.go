package relation

import (
	"context"
	"fmt"
	"monster/meta"
	"monster/system/datatype"
	"monster/system/datatype/composite"
	"monster/system/property"
	"monster/system/util"
)

// buildToManyExec build one-to-many relationshi
type buildToManyExec struct {
	Relation *composite.RelationType
	Target   meta.Instance

	capacity      int
	relationField meta.Field
}

func (exec *buildToManyExec) Open() (err error) {
	// read out relation capacity
	limit := exec.Relation.PropertyValue(property.RelationCapacity)
	if limit != nil {
		exec.capacity = limit.(int)
	} else {
		exec.capacity = 0
	}

	exec.relationField = util.RelationIndicatingField(exec.Relation)
	return
}

func (exec *buildToManyExec) Next(ctx context.Context, instances *meta.Batch) (err error) {
	errors := meta.AggregateError{}
	for _, ins := range instances.Items {
		// checkout capacity
		fieldValue := ins.FieldValue(exec.relationField)
		if fieldValue != nil && exec.capacity > 0 {
			targets := fieldValue.(meta.InstanceCollector)
			if targets.Size() >= exec.capacity {
				errors = append(errors, fmt.Errorf("cannot exceeds relation capacity %d", exec.capacity))
				continue
			}
		}
		// create relation instance
		newRelationInstance := composite.BuildRelationship(ins, exec.Target, exec.Relation)
		// copy all old relation instances
		collection := datatype.NewInstanceHashSet()
		if fieldValue != nil {
			for _, relationIns := range fieldValue.(meta.InstanceCollector).Values() {
				collection.Add(relationIns)
			}
		}
		// add newly created relation instance
		collection.Add(newRelationInstance)
		// set relation field value
		ins.SetFieldValue(exec.relationField, collection)
	}
	if len(errors) > 0 {
		err = errors
	}
	return
}

func (*buildToManyExec) Close() (err error) {
	return
}
