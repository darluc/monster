package meta

import "reflect"

type PropertyDefinition struct {
	name     string
	dataType reflect.Type
}

func (p *PropertyDefinition) Name() string {
	return p.name
}

func (p *PropertyDefinition) DataType() reflect.Type {
	return p.dataType
}

func NewPropertyDefinition(name string, t reflect.Type) *PropertyDefinition {
	return &PropertyDefinition{name: name, dataType: t}
}

// PropertyHolder have a batch of properties.
type PropertyHolder interface {
	Properties() []Property
	AddProperty(property Property)
}

// Property is often used by executors.
type Property interface {
	// Holder return the owner who has the property
	Holder() PropertyHolder
	SetHolder(holder PropertyHolder)
	Definition() *PropertyDefinition
	Value() interface{}
}
