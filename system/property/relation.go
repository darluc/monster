package property

import (
	"reflect"
	"monster/meta"
)

type Type uint8

const (
	TypeMasterSlave Type = iota
	TypeMasterReference
)

type Ratio uint8

const (
	OneToOne Ratio = iota
	OneToMany
)

var (
	RelationRatio = meta.NewPropertyDefinition("RelationType", reflect.TypeOf(TypeMasterSlave))
	RelationType  = meta.NewPropertyDefinition("RelationRatio", reflect.TypeOf(OneToOne))
)
