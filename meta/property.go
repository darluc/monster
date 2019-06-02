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

// Property is often used by executors.
type Property interface {
	// Holder return the owner who has the property
	Holder() PropertyHolder
	SetHolder(holder PropertyHolder)
	Definition() *PropertyDefinition
	Value() interface{}
}

// PropertyHolder have a batch of properties.
type PropertyHolder interface {
	Properties() []Property
	AddProperty(property Property)
	PropertyValue(definition *PropertyDefinition) interface{}
}

type BasePropertyHolder struct {
	properties []Property
}

func (base *BasePropertyHolder) Properties() []Property {
	return base.properties
}

func (base *BasePropertyHolder) AddProperty(property Property) {
	base.properties = append(base.properties, property)
}

func (base *BasePropertyHolder) PropertyValue(definition *PropertyDefinition) interface{} {
	for _, prop := range base.properties {
		if prop.Definition() == definition {
			return prop.Value()
		}
	}
	return nil
}

func NewProperties() *BasePropertyHolder {
	return &BasePropertyHolder{properties: make([]Property, 0)}
}
