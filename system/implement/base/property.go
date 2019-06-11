package base

import "monster/meta"

type Property struct {
	holder             meta.PropertyHolder
	propertyDefinition *meta.PropertyDefinition
	propertyValue      interface{}
}

func (p *Property) SetHolder(holder meta.PropertyHolder) {
	p.holder = holder
}

func (p *Property) Holder() meta.PropertyHolder {
	return p.holder
}

func (p *Property) Definition() *meta.PropertyDefinition {
	return p.propertyDefinition
}

func (p *Property) Value() interface{} {
	return p.propertyValue
}

func NewBaseProperty(definition *meta.PropertyDefinition, value interface{}) meta.Property {
	// @todo check value with definition's type
	return &Property{propertyDefinition: definition, propertyValue: value}
}
