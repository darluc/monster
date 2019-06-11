package base

import (
	"sync"
	"monster/meta"
)

// BaseObject is a simple implementation of Object
type Object struct {
	name   string
	fields []meta.Field
	*meta.BasePropertyHolder

	opMutex sync.Mutex
}

func (obj *Object) HasField(field meta.Field) bool {
	for _, f := range obj.fields {
		if f == field {
			return true
		}
	}
	return false
}

func (obj *Object) Name() string {
	return obj.name
}

func (obj *Object) Fields() []meta.Field {
	return obj.fields
}

func (obj *Object) Field(fieldName string) meta.Field {
	for _, f := range obj.fields {
		if f.Name() == fieldName {
			return f
		}
	}
	return nil
}

func (obj *Object) AddField(field meta.Field) {
	obj.opMutex.Lock()
	defer obj.opMutex.Unlock()
	obj.fields = append(obj.fields, field)
}

func NewBaseObject() meta.Object {
	b := &Object{}
	b.BasePropertyHolder = meta.NewProperties()
	b.fields = make([]meta.Field, 0)
	return b
}
