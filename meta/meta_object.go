package meta

// Object defines interface for meta object
type Object interface {
	Name() string
	Fields() []Field
	Field(fieldName string) Field
	AddField(field Field)
	HasField(field Field) bool
	PropertyHolder
}
