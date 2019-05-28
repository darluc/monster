package base

import (
	"errors"
	"fmt"
)

type Instance struct {
	metaObject Object
	fields     map[Field]interface{}
}

func (ins *Instance) SetFieldValue(field Field, value interface{}) {
	if !ins.metaObject.HasField(field) {
		panic(fmt.Errorf("object[%s] does not have field[%s]", ins.metaObject.Name(), field.Name()))
	}
	ins.fields[field] = value
}

func (ins *Instance) ID() Identifier {
	panic("implement me")
}

func (ins *Instance) MetaObject() Object {
	return ins.metaObject
}

func (ins *Instance) FieldValue(field Field) interface{} {
	if v, ok := ins.fields[field]; ok {
		return v
	}
	return nil
}

func NewBaseInstance(object Object) Instance {
	if object == nil {
		panic(errors.New("cannot create instance from a nil object"))
	}
	b := &Instance{fields: make(map[Field]interface{})}
	b.metaObject = object
	return b
}