package property

import (
	"monster/meta"
	"reflect"
)

type Type int

const (
	TypeMasterSlave Type = iota
	TypeMasterReference
)

type ratio int

const (
	OneToOne ratio = iota
	OneToMany
)

var (
	RelationRatio     = meta.NewPropertyDefinition("Relation:Ratio", reflect.TypeOf(TypeMasterSlave))
	RelationType      = meta.NewPropertyDefinition("Relation:Type", reflect.TypeOf(OneToOne))
	RelationFieldName = meta.NewPropertyDefinition("Relation:FieldNameOfSourceObject", reflect.TypeOf(""))
)
