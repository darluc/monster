package util

import (
	"theMoon/system/datatype"
	"time"
)

func GetCurrentDate() datatype.EffectiveDate {
	now := time.Now()
	y, mon, d := now.Date()
	// 20190601
	timeStr := y*100*100 + int(mon)*100 + d
	return datatype.EffectiveDate(timeStr)
}
