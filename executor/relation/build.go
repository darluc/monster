package relation

import (
	"context"
	"fmt"
	"monster/executor"
	"monster/meta"
	"monster/system/datatype/composite"
	"monster/system/property"
)

type BuildExec struct {
	Relation *composite.MetaDrivenType
	Target   meta.Instance

	underlyingExec executor.Executor // buildToOneExec / buildToManyExec
}

func (exec *BuildExec) Open() (err error) {
	if exec.Target == nil || exec.Relation == nil {
		return fmt.Errorf("relation target/type is nil")
	}
	if exec.Relation.PropertyValue(property.RelationCardinality) == property.OneToOne {
		exec.underlyingExec = &buildToOneExec{Relation: exec.Relation, Target: exec.Target}
	} else {
		exec.underlyingExec = &buildToManyExec{Relation: exec.Relation, Target: exec.Target}
	}
	return exec.underlyingExec.Open()
}

func (exec *BuildExec) Next(ctx context.Context, instances *meta.Batch) error {
	return exec.underlyingExec.Next(ctx, instances)
}

func (exec *BuildExec) Close() error {
	return exec.underlyingExec.Close()
}
