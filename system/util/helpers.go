package util

import (
	"monster/meta"
	"monster/system/datatype"
	"monster/system/datatype/composite"
	"monster/system/property"
	"time"
)

func GetCurrentDate() datatype.EffectiveDate {
	now := time.Now()
	y, mon, d := now.Date()
	// 20190601
	timeStr := y*100*100 + int(mon)*100 + d
	return datatype.EffectiveDate(timeStr)
}

// RelationIndicatingField extract the meta field which holds the relation data
func RelationIndicatingField(relationType *composite.MetaDrivenType) meta.Field {
	sourceObject := relationType.Field(composite.RelationSource).Type().(*composite.MetaDrivenType).Object
	return sourceObject.Field(relationType.PropertyValue(property.RelationIndicatingName).(string))
}
