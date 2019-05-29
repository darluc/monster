package timed

import (
	"theMoon/meta"
	"theMoon/system/datatype"
	"theMoon/system/implement/base"
	"theMoon/system/property"
)

type Object struct {
	*base.Object
}

func NewTimedObject() meta.Object {
	obj := &Object{Object: base.NewBaseObject().(*base.Object)}

	// set object timeline property to true
	enableTimeline := base.NewBaseProperty(property.TimelineEnabled, true)
	obj.AddProperty(enableTimeline)

	// add effectiveStartDate & effectiveEndDate field
	startDateField := base.NewBaseField(FieldEffectiveStartDate, datatype.EffectiveDateType)
	obj.AddField(startDateField)
	endDateField := base.NewBaseField(FieldEffectiveEndDate, datatype.EffectiveDateType)
	obj.AddField(endDateField)

	return obj
}
