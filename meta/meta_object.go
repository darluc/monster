package meta

import (
	"sync"
)

// Object defines interface for meta object
type Object interface {
	Name() string
	Fields() []Field
	Field(fieldName string) Field
	AddField(field Field)
	RelationWith(object Object) Object //@todo the datatype of return value not defined
	HasField(field Field) bool
	PropertyHolder
}

type ObjectConstructor func() Object

// BaseObject is a simple implementation of Object
type BaseObject struct {
	name       string
	fields     []Field
	properties []Property

	opMutex sync.Mutex
}

func (obj *BaseObject) RelationWith(object Object) Object {
	panic("implement me")
}

func (obj *BaseObject) HasField(field Field) bool {
	for _, f := range obj.fields {
		if f == field {
			return true
		}
	}
	return false
}

func (obj *BaseObject) Name() string {
	return obj.name
}

func (obj *BaseObject) Fields() []Field {
	return obj.fields
}

func (obj *BaseObject) Field(fieldName string) Field {
	for _, f := range obj.fields {
		if f.Name() == fieldName {
			return f
		}
	}
	return nil
}

func (obj *BaseObject) AddField(field Field) {
	obj.opMutex.Lock()
	defer obj.opMutex.Unlock()
	obj.fields = append(obj.fields, field)
}

func (obj *BaseObject) Properties() []Property {
	return obj.properties
}

func (obj *BaseObject) AddProperty(property Property) {
	obj.opMutex.Lock()
	defer obj.opMutex.Unlock()
	obj.properties = append(obj.properties, property)
}

func NewBaseObject() Object {
	b := &BaseObject{}
	b.properties = make([]Property, 0)
	b.fields = make([]Field, 0)
	return b
}
