package meta

type ObjectConstructor func(name string) Object
type FieldConstructor func(name string) Field
type PropertyConstructor func() Property
type InstanceConstructor func(object Object) Instance

type ConstructionSuite interface {
	ObjectConstructor() ObjectConstructor
	FieldConstructor() FieldConstructor
	InstanceConstructor() InstanceConstructor
	PropertyConstructor() PropertyConstructor
}
