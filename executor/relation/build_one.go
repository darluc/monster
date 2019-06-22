package relation

import (
	"context"
	"fmt"
	"monster/meta"
	"monster/system/datatype/composite"
	"monster/system/property"
	"monster/system/util"
)

// buildToOneExec build relation for one-to-one relationship
type buildToOneExec struct {
	Relation *composite.RelationType
	Target   meta.Instance

	garbage       *DestroyExec
	relationField meta.Field
}

func (exec *buildToOneExec) Open() (err error) {
	if exec.Relation.PropertyValue(property.RelationCardinality) != property.OneToOne {
		return fmt.Errorf("only OneToOne relation is allowed")
	}

	exec.garbage = &DestroyExec{Relation: exec.Relation}
	err = exec.garbage.Open()

	exec.relationField = util.RelationIndicatingField(exec.Relation)
	return
}

func (exec *buildToOneExec) Next(ctx context.Context, instances *meta.Batch) error {
	errors := meta.AggregateError{}
	for _, ins := range instances.Items {
		relationTarget := ins.FieldValue(exec.relationField)

		// recycling old relation target
		if relationTarget != nil && relationTarget != exec.Target {
			err := exec.garbage.Next(ctx, &meta.Batch{Items: []meta.Instance{ins}})
			if err != nil {
				errors = append(errors, err)
				continue
			}
		}

		// create new relation instance
		newRelationTarget := composite.BuildRelationship(ins, exec.Target, exec.Relation)
		ins.SetFieldValue(exec.relationField, newRelationTarget)
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

func (exec *buildToOneExec) Close() (err error) {
	return exec.garbage.Close()
}
