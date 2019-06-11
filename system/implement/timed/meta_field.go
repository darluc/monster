package timed

import (
	"monster/meta"
	"monster/system/implement/base"
	"monster/system/property"
)

const (
	FieldEffectiveStartDate string = "effectiveStartDate"
	FieldEffectiveEndDate   string = "effectiveEndDate"
	FieldEffectiveEntryTime string = "effectiveEntryTime"
)

type Field struct {
	*base.Field
}

func NewTimedField(name string, dataType meta.DataType) meta.Field {
	field := Field{Field: base.NewBaseField(name, dataType).(*base.Field)}
	// set object timeline property to true
	enableTimeline := base.NewBaseProperty(property.TimelineEnabled, true)
	field.AddProperty(enableTimeline)
	return field
}
