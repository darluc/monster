package timed

import (
	"errors"
	"github.com/emirpasic/gods/maps/treemap"
	"monster/meta"
)

type TimeCapsule struct {
	values *treemap.Map
}

// SetCurrentValue sets value at current view of time.
func (t *TimeCapsule) SetCurrentValue(val interface{}) {
	ct := meta.NewEntryTime(false)
	lastTime, _ := t.latestEntryTime(ct)
	if lastTime.IsEmpty() {
		t.values.Put(meta.GenesisEntryTime, val)
	} else {
		t.values.Put(lastTime, val)
	}
	return
}

// latestEntryTime search for the entry time & value can be seen at the search time.
func (t *TimeCapsule) latestEntryTime(searchTime meta.EntryTime) (latestTime meta.EntryTime, latestValue interface{}) {
	var lastEntryTime meta.EntryTime
	latestValue, found := t.values.Get(searchTime)
	if found {
		lastEntryTime = searchTime
		return
	}

	var lastValue interface{}
	it := t.values.Iterator()
	for it.Next() {
		k, v := it.Key(), it.Value()
		et := k.(meta.EntryTime)
		if searchTime.LaterThan(et) {
			lastEntryTime = et
			lastValue = v
			continue
		} else {
			latestTime = lastEntryTime
			latestValue = lastValue
			return
		}
	}
	if !lastEntryTime.IsEmpty() {
		latestTime = lastEntryTime
		latestValue = lastValue
	} else {
		latestValue = nil
	}
	return
}

// ValueAt gets value at specified entry time. It will return nil if no value found.
func (t *TimeCapsule) ValueAt(time meta.EntryTime) interface{} {
	_, v := t.latestEntryTime(time)
	return v
}

// SetValueAt set value at the specified entry time. If the action is not forced and the entry time does not
// exist, it will return an error.
func (t *TimeCapsule) SetValueAt(val interface{}, time meta.EntryTime, force bool) (err error) {
	if _, found := t.values.Get(time); found && !force {
		err = errors.New("entry time collision")
	}
	t.values.Put(time, val)
	return
}

// Slices return an time iterator for values of all time slices.
func (t *TimeCapsule) Slices() meta.TimeIterator {
	return newTimeCapsuleIterator(t)
}

type TimeCapsuleIterator struct {
	treeIterator treemap.Iterator
}

func (t *TimeCapsuleIterator) Next() bool {
	return t.treeIterator.Next()
}

func (t *TimeCapsuleIterator) EntryTime() meta.EntryTime {
	return t.treeIterator.Key().(meta.EntryTime)
}

func (t *TimeCapsuleIterator) Value() interface{} {
	return t.treeIterator.Value()
}

func newTimeCapsuleIterator(capsule *TimeCapsule) *TimeCapsuleIterator {
	return &TimeCapsuleIterator{treeIterator: capsule.values.Iterator()}
}

func entryTimeComparator(a interface{}, b interface{}) int {
	at := a.(meta.EntryTime)
	bt := b.(meta.EntryTime)
	if at.Date == bt.Date {
		if at.Time == bt.Time {
			return 0
		} else {
			if at.Time < bt.Time {
				return -1
			} else {
				return 1
			}
		}
	} else {
		if at.Date < bt.Date {
			return -1
		} else {
			return 1
		}
	}
}

func NewTimeCapsule() *TimeCapsule {
	t := new(TimeCapsule)
	t.values = treemap.NewWith(entryTimeComparator)
	return t
}
