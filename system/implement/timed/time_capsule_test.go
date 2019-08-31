package timed

import (
	"monster/meta"
	"testing"
)

func TestTimeCapsule_SetCurrentValue(t *testing.T) {
	tc := NewTimeCapsule()
	tc.SetCurrentValue(555)
	tc.SetValueAt(666, meta.EntryTime{Date: 19900101}, true)

	expect := 666
	if cv := tc.ValueAt(meta.NewEntryTime(false)); cv == expect {
		t.Logf("current value is %d", expect)
	} else {
		t.Errorf("current value is not %d but %d", expect, cv)
	}

	tc1 := NewTimeCapsule()
	tc1.SetValueAt(666, meta.EntryTime{Date: 19900101}, true)
	tc1.SetCurrentValue(555)

	expect = 555
	if cv := tc1.ValueAt(meta.NewEntryTime(false)); cv == expect {
		t.Logf("current value is %d", expect)
	} else {
		t.Errorf("current value is not %d but %d", expect, cv)
	}
}

func TestTimeCapsule_SetValueAt(t *testing.T) {
	tc := NewTimeCapsule()
	dataEntryTime := meta.EntryTime{
		Date: 19000101,
		Time: 10,
	}
	currentEntryTime := meta.NewEntryTime(false)
	val := 555
	tc.SetValueAt(val, dataEntryTime, false)
	if v := tc.ValueAt(currentEntryTime); v != val {
		t.Errorf("current value is %d and should be %d", v, val)
	} else {
		t.Logf("value is %d", v)
	}

	newVal := 666
	if err := tc.SetValueAt(newVal, dataEntryTime, false); err == nil {
		t.Errorf("should get error without enforcement")
	} else {
		t.Logf("set value with same entrytime is not allowed")
	}

	tc.SetValueAt(newVal, dataEntryTime, true)
	if v := tc.ValueAt(currentEntryTime); v != newVal {
		t.Errorf("new value is %d and should be %d", v, newVal)
	} else {
		t.Logf("new value is %d", v)
	}
}

func TestTimeCapsule_Slices(t *testing.T) {
	tc := NewTimeCapsule()
	tc.SetValueAt(2, meta.EntryTime{Date: 19000101, Time: 0,}, false)
	tc.SetValueAt(3, meta.EntryTime{Date: 19000101, Time: 10,}, false)
	tc.SetValueAt(0, meta.EntryTime{Date: 17000101, Time: 10,}, false)
	tc.SetValueAt(4, meta.EntryTime{Date: 20000101, Time: 0,}, false)
	tc.SetValueAt(1, meta.EntryTime{Date: 18000101, Time: 0,}, false)

	it := tc.Slices()
	last := -1
	for it.Next() {
		if it.Value().(int) < last {
			t.Errorf("value is bigger than last value")
			return
		} else {
			t.Logf("value %d:%d is %d", it.EntryTime().Date, it.EntryTime().Time, it.Value())
		}
		last = it.Value().(int)
	}
	t.Logf("all values are sorted by entry time")
}
