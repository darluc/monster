package executor

import (
	"context"
	"errors"
	"monster/meta"
)

type CreateExec struct {
	InstanceCreator meta.InstanceConstructor
	MetaObject      meta.Object
	Data            map[string]interface{}

	update *UpdateExec
}

func (e *CreateExec) Open() error {
	if e.MetaObject == nil {
		return errors.New("meta object is nil")
	}
	if e.Data != nil && len(e.Data) > 0 {
		e.update = &UpdateExec{Data: e.Data}
		return e.update.Open()
	}
	return nil
}

func (e *CreateExec) Next(ctx context.Context, instances *meta.Batch) (err error) {
	ins := e.InstanceCreator(e.MetaObject)
	instances.Items = append(instances.Items, ins)
	if e.update != nil {
		err = e.update.Next(ctx, instances)
	}
	return err
}

func (e *CreateExec) Close() error {
	if e.update != nil {
		return e.update.Close()
	}
	return nil
}
