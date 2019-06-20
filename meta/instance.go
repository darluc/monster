package meta

import (
	"fmt"
)

type Identifier interface {
	Equals(identifier Identifier) bool
	fmt.Stringer
}

// Instance defines interface for instance datatype
type Instance interface {
	ID() Identifier
	// MetaObject returns the meta object of the instance
	MetaObject() Object
	// FieldValue help you get the actual value of the specified meta field
	FieldValue(field Field) interface{}
	// SetFieldValue set value for field
	SetFieldValue(field Field, value interface{})
}

// Batch is used to batch operation, especially for executors
type Batch struct {
	Items []Instance
}

func NewBatch() *Batch {
	b := &Batch{}
	b.Items = make([]Instance, 0)
	return b
}
