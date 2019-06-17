package executor

import (
	"context"
	"errors"
	"monster/meta"
)

// CreateExec is a basic executor used to create instances
type CreateExec struct {
	InstanceCreator meta.InstanceConstructor
	MetaObject      meta.Object
}

func (e *CreateExec) Open() (err error) {
	if e.MetaObject == nil {
		err = errors.New("meta object cannot be nil")
	}
	return
}

func (e *CreateExec) Next(ctx context.Context, instances *meta.Batch) (err error) {
	ins := e.InstanceCreator(e.MetaObject)
	instances.Items = append(instances.Items, ins)
	return err
}

func (e *CreateExec) Close() error {
	return nil
}
