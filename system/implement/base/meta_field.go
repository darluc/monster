package base

import (
	"reflect"
	"sync"
	"theMoon/meta"
)

type Field struct {
	name       string
	dataType   reflect.Type
	properties []meta.Property
	opMutex    sync.Mutex
}

func NewBaseField(name string, t reflect.Type) *Field {
	b := &Field{name: name, dataType: t}
	b.properties = make([]meta.Property, 0)
	return b
}

func (field *Field) Name() string {
	return field.name
}

func (field *Field) Type() reflect.Type {
	return field.dataType
}

func (field *Field) SetName(name string) {
	field.name = name
}

func (field *Field) Properties() []meta.Property {
	return field.properties
}

func (field *Field) addProperty(property meta.Property) {
	field.opMutex.Lock()
	defer field.opMutex.Unlock()
	field.properties = append(field.properties, property)
}
