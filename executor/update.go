package executor

import (
	"context"
	"errors"
	"monster/meta"
)

// UpdateExec is a basic executor used to update instance field's value
type UpdateExec struct {
	Data map[string]interface{}
}

func (e *UpdateExec) Open() error {
	if len(e.Data) > 0 {
		return errors.New("empty datatype set")
	}
	return nil
}

func (e *UpdateExec) Next(ctx context.Context, instances *meta.Batch) error {
	for _, ins := range instances.Items {
		metaObj := ins.MetaObject()
		for _, fld := range metaObj.Fields() {
			if v, ok := e.Data[fld.Name()]; ok {
				ins.SetFieldValue(fld, v)
			}
		}
	}
	return nil
}

func (e *UpdateExec) Close() error {
	return nil
}
