package timed

import "monster/system/datatype"

type timedValue struct {
	EntryTime uint
	Date      datatype.EffectiveDate
	Value     interface{}
}
