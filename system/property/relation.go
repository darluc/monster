package property

import (
	"monster/meta"
	"reflect"
)

type tRelationCoupling int

const (
	MasterSlave tRelationCoupling = iota
	MasterReference
)

type TRelationCardinality int

const (
	OneToOne TRelationCardinality = iota
	OneToMany
)

var (
	RelationCapacity       = meta.NewPropertyDefinition("Relation:Capacity", reflect.TypeOf(1))
	RelationCardinality    = meta.NewPropertyDefinition("Relation:Cardinality", reflect.TypeOf(OneToOne))
	RelationCoupling       = meta.NewPropertyDefinition("Relation:Category", reflect.TypeOf(MasterSlave))
	RelationIndicatingName = meta.NewPropertyDefinition("Relation:IndicatingName", reflect.TypeOf(""))
)
