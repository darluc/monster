package meta

import "time"

var GenesisEntryTime = EntryTime{Date: 10000101, Time: 0}

type EntryTime struct {
	Date uint
	Time uint
}

func (t EntryTime) IsEmpty() bool {
	return t.Date == 0 && t.Time == 0
}

func (t EntryTime) LaterThan(ct EntryTime) bool {
	if t.Date == ct.Date {
		return t.Time > ct.Time
	} else {
		return t.Date > ct.Date
	}
}

func (t EntryTime) Equal(ct EntryTime) bool {
	return t.Date == ct.Date && t.Time == ct.Time
}

func NewEntryTime(withUnixTime bool) EntryTime {
	now := time.Now()
	y, mon, d := now.Date()
	dateInt := y*100*100 + int(mon)*100 + d
	et := EntryTime{}
	et.Date = uint(dateInt)
	if withUnixTime {
		et.Time = uint(now.Unix())
	} else {
		et.Time = 0
	}
	return et
}

type (
	TimeContainer interface {
		ValueAt(time EntryTime) interface{}
		SetValueAt(val interface{}, time EntryTime, force bool) error
		// setCurrentValue changes the data value in the slice of current view
		SetCurrentValue(val interface{})
		Slices() TimeIterator
	}

	TimeIterator interface {
		Next() bool
		EntryTime() EntryTime
		Value() interface{}
	}
)
