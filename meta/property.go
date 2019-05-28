package meta

type PropertyName string

// PropertyHolder have a batch of properties.
type PropertyHolder interface {
	Properties() []Property
}

// Property is often used by executors.
type Property interface {
	// Holder return the owner who has the property
	Holder() PropertyHolder
	Name() PropertyName
	Value() interface{}
}
