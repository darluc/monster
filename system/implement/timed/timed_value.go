package timed

import "theMoon/system/datatype"

type timedValue struct {
	EntryTime uint
	Date      datatype.EffectiveDate
	Value     interface{}
}
