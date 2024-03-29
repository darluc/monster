package base

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"monster/meta"
)

type Instance struct {
	metaObject meta.Object
	fields     map[meta.Field]interface{}
}

func (ins *Instance) SetFieldValue(field meta.Field, value interface{}) {
	if !ins.metaObject.HasField(field) {
		panic(fmt.Errorf("object[%s] does not have field[%s]", ins.metaObject.Name(), field.Name()))
	}
	if field.Type().TypeCheck(value) {
		ins.fields[field] = value
	} else {
		logrus.Warnf("set field[%s] with improper value[%v]", field.Name(), value)
	}
}

func (ins *Instance) ID() meta.Identifier {
	panic("implement me")
}

func (ins *Instance) MetaObject() meta.Object {
	return ins.metaObject
}

func (ins *Instance) FieldValue(field meta.Field) interface{} {
	if v, ok := ins.fields[field]; ok {
		return v
	}
	return nil
}

func (ins *Instance) FieldValueByName(fieldName string) interface{} {
	if fld := ins.MetaObject().Field(fieldName); fld != nil {
		return ins.FieldValue(fld)
	}
	return nil
}

func NewBaseInstance(object meta.Object) meta.Instance {
	if object == nil {
		panic(errors.New("cannot create instance from a nil object"))
	}
	b := &Instance{fields: make(map[meta.Field]interface{})}
	b.metaObject = object
	return b
}
