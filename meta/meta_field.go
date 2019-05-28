package meta

// Field defines basic interface for field of meta object
type Field interface {
	Name() string
	Type() DataType
	PropertyHolder
}
