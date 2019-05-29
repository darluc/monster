package base

import (
	"sync"
	"theMoon/meta"
)

type Field struct {
	name       string
	dataType   meta.DataType
	properties []meta.Property
	opMutex    sync.Mutex
}

func (field *Field) Name() string {
	return field.name
}

func (field *Field) Type() meta.DataType {
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

func NewBaseField(name string, t meta.DataType) meta.Field {
	b := &Field{name: name, dataType: t}
	b.properties = make([]meta.Property, 0)
	return b
}
