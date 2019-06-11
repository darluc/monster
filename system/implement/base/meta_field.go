package base

import (
	"sync"
	"monster/meta"
)

type Field struct {
	name     string
	dataType meta.DataType
	opMutex  sync.Mutex

	*meta.BasePropertyHolder
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

func NewBaseField(name string, t meta.DataType) meta.Field {
	b := &Field{name: name, dataType: t, BasePropertyHolder: meta.NewProperties()}
	return b
}
