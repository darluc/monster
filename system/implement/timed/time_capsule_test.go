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
