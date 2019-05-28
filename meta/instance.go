package meta

import (
	"errors"
	"fmt"
)

type Identifier interface {
	Equals(identifier Identifier) bool
	fmt.Stringer
}

// Instance defines interface for instance datatype
type Instance interface {
	ID() Identifier
	// MetaObject returns the meta object of the instance
	MetaObject() Object
	// FieldValue help you get the actual value of the specified meta field
	FieldValue(field Field) interface{}
	// SetFieldValue set value for field
	SetFieldValue(field Field, value interface{})
}

type Batch struct {
	Items []Instance
}

func NewBatch() *Batch {
	b := &Batch{}
	b.Items = make([]Instance, 0)
	return b
}

// Constructor is a function used to create instance
type InstanceConstructor func(object Object) Instance

type BaseInstance struct {
	metaObject Object
	fields     map[Field]interface{}
}

func (ins *BaseInstance) SetFieldValue(field Field, value interface{}) {
	if !ins.metaObject.HasField(field) {
		panic(fmt.Errorf("object[%s] does not have field[%s]", ins.metaObject.Name(), field.Name()))
	}
	ins.fields[field] = value
}

func (ins *BaseInstance) ID() Identifier {
	panic("implement me")
}

func (ins *BaseInstance) MetaObject() Object {
	return ins.metaObject
}

func (ins *BaseInstance) FieldValue(field Field) interface{} {
	if v, ok := ins.fields[field]; ok {
		return v
	}
	return nil
}

func NewBaseInstance(object Object) Instance {
	if object == nil {
		panic(errors.New("cannot create instance from a nil object"))
	}
	b := &BaseInstance{fields: make(map[Field]interface{})}
	b.metaObject = object
	return b
}
