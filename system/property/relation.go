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
	RelationRatio = meta.NewPropertyDefinition("RelationRatio", reflect.TypeOf(TypeMasterSlave))
	RelationType  = meta.NewPropertyDefinition("RelationType", reflect.TypeOf(OneToOne))
)
