package property

import (
	"reflect"
	"theMoon/meta"
)

var (
	EffectiveStartDate = meta.NewPropertyDefinition("EffectiveStartDate", reflect.TypeOf(uint(1)))
	EffectiveEndDate   = meta.NewPropertyDefinition("EffectiveStartDate", reflect.TypeOf(uint(1)))
)

type TimelineSupported bool

var (
	TimelineEnabled = meta.NewPropertyDefinition("EffectiveStartDate", reflect.TypeOf(true))
)
