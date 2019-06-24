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

// ExtendsObject adds fields of other objects to target object
func ExtendsObject(target Object, objects ...Object) Object {
	for _, obj := range objects {
		for _, fld := range obj.Fields() {
			target.AddField(fld)
		}
	}
	return target
}
