package property

import (
	"reflect"
	"theMoon/meta"
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
	RelationRatio *meta.PropertyDefinition
	RelationType  *meta.PropertyDefinition
)

func init() {
	RelationType = meta.NewPropertyDefinition("RelationType", reflect.TypeOf(TypeMasterSlave))
	RelationRatio = meta.NewPropertyDefinition("RelationRatio", reflect.TypeOf(OneToOne))
}
