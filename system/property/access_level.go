package property

import (
	"monster/meta"
	"reflect"
)

type tAccessLevel int

const (
	PublicAccess tAccessLevel = iota
	InnerAccess
)

var (
	FieldAccessibility = meta.NewPropertyDefinition("Field:Accessibility", reflect.TypeOf(1))
)
