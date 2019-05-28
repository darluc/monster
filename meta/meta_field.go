package meta

import (
	"reflect"
	"sync"
)

// Field defines basic interface for field of meta object
type Field interface {
	Name() string
	Type() reflect.Type
	PropertyHolder
}

type FieldConstructor func(name string) Field

// BaseField is a simple implementation of Field
type BaseField struct {
	name       string
	dataType   reflect.Type
	properties []Property
	opMutex    sync.Mutex
}

func NewBaseField(name string, t reflect.Type) *BaseField {
	b := &BaseField{name: name, dataType: t}
	b.properties = make([]Property, 0)
	return b
}

func (field *BaseField) Name() string {
	return field.name
}

func (field *BaseField) Type() reflect.Type {
	return field.dataType
}

func (field *BaseField) SetName(name string) {
	field.name = name
}

func (field *BaseField) Properties() []Property {
	return field.properties
}

func (field *BaseField) addProperty(property Property) {
	field.opMutex.Lock()
	defer field.opMutex.Unlock()
	field.properties = append(field.properties, property)
}
